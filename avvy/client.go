package avvy


import (
   "log"
   "os"
   "io/ioutil"
   "math/big"
   "strings"
   "strconv"
   "path"
   "encoding/json"
   "github.com/ethereum/go-ethereum/common"
   "github.com/ethereum/go-ethereum/ethclient"
   poseidon "github.com/avvydomains/golang-client/abi/Poseidon"
)

type ClientCommonContract struct {
    Address string `json:"address"`
}

type ClientCommon struct {
    Contracts map[string](ClientCommonContract) `json:"contracts"`
}

type Client struct {
    rpcUrl string
    clientCommonPath string
    clientCommon ClientCommon
    chainId int
    client *ethclient.Client
}

/* Loads in relevant data from Client Common lib */
func (c *Client) InitClientCommon() {
    contractPath := path.Join(c.clientCommonPath, "contracts", strconv.Itoa(c.chainId) + ".json")
    content, err := ioutil.ReadFile(contractPath)
    if err != nil {
        log.Fatal("Error opening contracts file from client common", err)
    }
    var res ClientCommon
    err = json.Unmarshal(content, &res)
    if err != nil {
        log.Fatal("Error converting contracts file to JSON", err)
    }
    c.clientCommon = res
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
    return c.clientCommon.Contracts[contractName].Address
}

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
            l2[i] = uint64(char)
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

