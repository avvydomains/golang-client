package avvy


import (
    "log"
    "os"
    "math/big"
    "strings"
    "strconv"
    "embed"
    "encoding/json"
    "time"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    // contracts
    domain "github.com/avvydomains/golang-client/abi/Domain"
    poseidon "github.com/avvydomains/golang-client/abi/Poseidon"
    publicResolverV1 "github.com/avvydomains/golang-client/abi/PublicResolverV1"
    resolverRegistryV1 "github.com/avvydomains/golang-client/abi/ResolverRegistryV1"
)

//go:embed contracts.json
//go:embed records.json
var embeddedFiles embed.FS

type ClientCommonContract struct {
    Address string `json:"address"`
}

type ClientCommonContracts struct {
    Contracts map[string](ClientCommonContract) `json:"contracts"`
}

type ClientCommonContractsBase struct {
    Chains map[string](ClientCommonContracts) `json:"chains"`
}

type ClientCommonRecord struct {
    Key int `json:"key"`
    Name string `json:"name"`
    Label string `json:"label"`
    Description string `json:"description"`
}

type ClientCommonRecordsBase struct {
    Records []ClientCommonRecord `json:"records"`
}

type ResolverAddress struct {
    Resolver common.Address
    DatasetId *big.Int
}

type Client struct {
    rpcUrl string
    clientCommonPath string
    clientCommonContracts ClientCommonContracts
    chainId int
    client *ethclient.Client
    RECORDS map[string](big.Int)
}

/* Loads in relevant data from Client Common lib */
func (c *Client) InitClientCommon() {

    // init contracts
    content, err := embeddedFiles.ReadFile("contracts.json")
    if err != nil {
        log.Fatal("Error opening contracts file from client common", err)
    }
    var res ClientCommonContractsBase
    err = json.Unmarshal(content, &res)
    if err != nil {
        log.Fatal("Error converting contracts file to JSON", err)
    }
    chainIdStr := strconv.Itoa(c.chainId)
    c.clientCommonContracts = res.Chains[chainIdStr]

    // init records
    content, err = embeddedFiles.ReadFile("records.json")
    if err != nil {
        log.Fatal("Error opening records file from client common", err)
    }
    var res2 ClientCommonRecordsBase
    err = json.Unmarshal(content, &res2)
    if err != nil {
        log.Fatal("Error converting records file to JSON", err)
    }
    m := make(map[string]big.Int)
    for _, record := range res2.Records {
        m[record.Name] = *big.NewInt(int64(record.Key))
    }
    c.RECORDS = m
}

/* Initializes the client */
func (c *Client) Init(rpcUrl string, chainId int) {
    clientCommonPath, present := os.LookupEnv("AVVY_CLIENT_COMMON")
    if !present {
        clientCommonPath = "client-common"
    }
    c.clientCommonPath = clientCommonPath
    c.chainId = chainId
    c.InitClientCommon()

    c.rpcUrl = rpcUrl
    client, err := ethclient.Dial(c.rpcUrl)
    if err != nil {
        log.Fatal(err)
    }
    c.client = client
}

/* Fetches the address for a given contract */
func (c *Client) GetContractAddress(contractName string) string {
    return c.clientCommonContracts.Contracts[contractName].Address
}

/* Gets Poseidon hash of provided name */
func (c *Client) Poseidon(inputNums [3]big.Int) big.Int {
    poseidonAddress := c.GetContractAddress("Poseidon")
    address := common.HexToAddress(poseidonAddress)
    instance, err := poseidon.NewPoseidon(address, c.client)
    if err != nil {
        log.Fatal(err)
    }
    var inputBytes [3][32]byte
    for i, n := range inputNums {
        numBytes := n.Bytes()
        for j := 0; j < len(numBytes); j += 1 {
            inputBytes[i][32 - len(numBytes) + j] = numBytes[j]
        }
    }
    outputBytes, err := instance.Poseidon(nil, inputBytes)
    if err != nil {
        log.Fatal(err)
    }
    z := new(big.Int)
    z.SetBytes(outputBytes[0:32])
    return *z
}

/* Checks the expiry date for a given name */
func (c *Client) GetExpiry(hash big.Int) big.Int {
    address := common.HexToAddress(c.GetContractAddress("Domain"))
    instance, err := domain.NewDomain(address, c.client)
    if err != nil {
        log.Fatal("Failed to initialize contract")
    }
    expiry, err := instance.GetDomainExpiry(nil, &hash)
    return *expiry
}

/* Checks if a given hash is expired */
func (c *Client) IsExpired(hash big.Int) bool {
    expiry := c.GetExpiry(hash)
    now := time.Now()
    sec := big.NewInt(now.Unix())
    return expiry.CmpAbs(sec) == -1
}

/* Gets the resolver for a given hash */
func (c *Client) GetResolver(domain big.Int, hash big.Int) (ResolverAddress, bool) {
    address := common.HexToAddress(c.GetContractAddress("ResolverRegistryV1"))
    instance, err := resolverRegistryV1.NewResolverRegistryV1(address, c.client)
    if err != nil {
        log.Fatal("Failed to initialize Resolver Registry contract")
    }
    resolverAddress, err := instance.Get(nil, &domain, &hash)
    if err != nil {
        return ResolverAddress{}, false
    }
    return resolverAddress, true
}

/* Gets the domain and hash for a given name */
func (c *Client) GetDomainAndHash(name string) (big.Int, big.Int) {
    nameLower := strings.ToLower(name)
    split := strings.Split(nameLower, ".")
    numLabels := len(split)
    if numLabels < 2 {
        log.Fatal("GetDomainAndHash must be called with a string that has at least two labels, separated by a period (.)")
    }
    topLabel := split[numLabels - 1]
    secondLabel := split[numLabels - 2]
    tld := secondLabel + "." + topLabel
    domain := c.NameHash(tld)
    var hash big.Int
    if tld == nameLower {
        hash = domain
    } else {
        hash = c.NameHash(nameLower)
    }
    return domain, hash
}

/* Resolves a Standard Record for a given name */
func (c *Client) ResolveStandard(name string, key big.Int) (string, bool) {
    domain, hash := c.GetDomainAndHash(name)
    if c.IsExpired(hash) {
        return "", false
    }
    resolver, success := c.GetResolver(domain, hash)
    if !success {
        return "", false
    }
    instance, err := publicResolverV1.NewPublicResolverV1(resolver.Resolver, c.client)
    if err != nil {
        return "", false
    }
    record, err := instance.ResolveStandard(nil, resolver.DatasetId, &hash, &key)
    if err != nil {
        return "", false
    }
    return record, true
}

/* Resolves a Custom Record for a given name */
func (c *Client) Resolve(name string, key string) (string, bool) {
    domain, hash := c.GetDomainAndHash(name)
    if c.IsExpired(hash) {
        return "", false
    }
    resolver, success := c.GetResolver(domain, hash)
    if !success {
        return "", false
    }
    instance, err := publicResolverV1.NewPublicResolverV1(resolver.Resolver, c.client)
    if err != nil {
        return "", false
    }
    record, err := instance.Resolve(nil, resolver.DatasetId, &hash, key)
    if err != nil {
        return "", false
    }
    return record, true
}

func num2Bits(inputNum uint64, numBits uint64) []uint64 {
    var lc1 uint64 = 0
    var e2 uint64 = 1
    var i uint64 = 0
    out := make([]uint64, numBits)
    for i = 0; i < numBits - 1; i = i + 1 {
        out[i] = (inputNum >> i) & 1
        lc1 += out[i] * e2
        e2 = e2 + e2
    }
    return out
}

func bits2BigInt(inputBits []uint64) big.Int {
    lc1 := big.NewInt(0)
    e2 := big.NewInt(1)
    for i := 0; i < len(inputBits); i += 1 {
        bi := new(big.Int).SetUint64(inputBits[i])
        lc1 = lc1.Add(lc1, bi.Mul(bi, e2))
        e2 = e2.Add(e2, e2)
    }
    return *lc1
}

func asciiArray2PreimageSignal(input [31]uint64) big.Int {
    var bits []uint64 = make([]uint64, 31*8)
    for i, char := range input {
        charBits := num2Bits(char, 8)
        for j, c := range charBits {
            bits[i*8+j] = c
        }
    }
    return bits2BigInt(bits)
}

/* Executes a single iteration of hashing */
func (c *Client) NameHashIteration(prevHash big.Int, label string) big.Int {
    var l1 [31]uint64
    var l2 [31]uint64
    for i, char := range label {
        if i < 31 {
            l1[i] = uint64(char)
        } else {
            l2[i-31] = uint64(char)
        }
    }
    n1 := asciiArray2PreimageSignal(l1)
    n2 := asciiArray2PreimageSignal(l2)
    var inputs [3]big.Int
    inputs[0] = prevHash
    inputs[1] = n1
    inputs[2] = n2
    return c.Poseidon(inputs)
}

/* Hashes the provided name */
func (c *Client) NameHash(domain string) big.Int {
    labels := strings.Split(domain, ".")

    // reverse labels
    for i, j := 0, len(labels)-1; i < j; i, j = i+1, j-1 {
        labels[i], labels[j] = labels[j], labels[i]
    }

    hash := big.NewInt(0)
    for _, label := range labels {
        *hash = c.NameHashIteration(*hash, label)
    }

    return *hash
}

