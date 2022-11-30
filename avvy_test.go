package avvy

import (
    "testing"
    "fmt"
    "math/big"
    "github.com/avvydomains/golang-client/avvy"
)

func buildClient() *avvy.Client {
    client := new(avvy.Client)
    client.Init("http://localhost:8545", 31337)
    return client
}

func TestNameHash(t *testing.T) {
    client := buildClient()
    hash := client.NameHash("test.avax")
    var expected, _ = new(big.Int).SetString("14724902060486453653675032315574462175944350516339223392990438905292905511078", 0)
    if hash.CmpAbs(expected) != 0 {
        fmt.Println(hash)
        fmt.Println(expected)
        t.Fatalf("Test hash does not match")
    }
}

func TestLongHash(t *testing.T) {
    client := buildClient()
    hash := client.NameHash("ascension-and-tristan-da-cunha-saint-helena.avax")
    var expected, _ = new(big.Int).SetString("19633460972774841899738637682877666780787918294612213522272481894413590645285", 0)
    if hash.CmpAbs(expected) != 0 {
        fmt.Println(hash)
        fmt.Println(expected)
        t.Fatalf("Test hash does not match")
    }
}

func TestExpiry(t *testing.T) {
    client := buildClient()
    hash := client.NameHash("avvy-client-common-testing.avax")
    expiry := client.GetExpiry(hash)
    var lowerBound, _ = new(big.Int).SetString("1601306611", 0)
    var higherBound, _ = new(big.Int).SetString("9901306611", 0)
    if expiry.CmpAbs(lowerBound) != 1 {
        t.Fatalf("Expiry date should be larger")
    }
    if expiry.CmpAbs(higherBound) != -1 {
        t.Fatalf("Expiry date should be smaller")
    }
}

func TestIsExpiredNo(t *testing.T) {
    client := buildClient()
    hash := client.NameHash("avvy-client-common-testing.avax")
    expired := client.IsExpired(hash)
    if expired {
        t.Fatalf("Domain should not be expired")
    }
}

func TestIsExpiredYes(t *testing.T) {
    client := buildClient()
    hash := client.NameHash("avvy-client-common-expired.avax")
    expired := client.IsExpired(hash)
    if !expired {
        t.Fatalf("Domain should be expired")
    }
}

func TestIsExpiredNotRegistered(t *testing.T) {
    client := buildClient()
    hash := client.NameHash("avvy-client-common-not-registered.avax")
    expired := client.IsExpired(hash)
    if !expired {
        t.Fatalf("Domain should be expired")
    }
}

func TestResolveNotRegistered(t *testing.T) {
    client := buildClient()
    _, useful := client.Resolve("avvy-client-common-not-registered.avax", "test")
    if useful {
        t.Fatalf("Expired domain should not return value")
    }
}

func TestResolveExpired(t *testing.T) {
    client := buildClient()
    _, useful := client.Resolve("avvy-client-common-expired.avax", "test")
    if useful {
        t.Fatalf("Expired domain should not return value")
    }
}

func TestResolveStandardNotRegistered(t *testing.T) {
    client := buildClient()
    var key, _ = new(big.Int).SetString("1", 0)
    _, useful := client.ResolveStandard("avvy-client-common-not-registered.avax", *key)
    if useful {
        t.Fatalf("Expired domain should not return value")
    }
}

func TestResolveStandardExpired(t *testing.T) {
    client := buildClient()
    var key, _ = new(big.Int).SetString("1", 0)
    _, useful := client.ResolveStandard("avvy-client-common-testing-expired.avax", *key)
    if useful {
        t.Fatalf("Expired domain should not return value")
    }
}

func TestGetResolver(t *testing.T) {
    client := buildClient()
    hash := client.NameHash("avvy-client-common-testing.avax")
    address, _ := client.GetResolver(hash, hash)
    if (*address.DatasetId).CmpAbs(&hash) != 0 {
        t.Fatalf("DatasetId on resolver seems incorrect")
    }
}

func TestGetResolverNotSet(t *testing.T) {
    client := buildClient()
    var hash, _ = new(big.Int).SetString("1601306611", 0)
    _, success := client.GetResolver(*hash, *hash)

    if success {
        t.Fatalf("GetResolver should fail")
    }
}

func TestResolveCustom(t *testing.T) {
    client := buildClient()
    record, success := client.Resolve("avvy-client-common-testing.avax", "CUSTOM_KEY")
    if !success {
        t.Fatalf("Should succeed")
    }
    if record != "CUSTOM_VALUE" {
        t.Fatalf("Record is wrong")
    }
}

func TestResolveStandard(t *testing.T) {
    client := buildClient()
    record, success := client.ResolveStandard("avvy-client-common-testing.avax", client.RECORDS["X_CHAIN"])
    if !success {
        t.Fatalf("Should succeed")
    }
    if record != "x-avax13fd740ykwc5peewmkcgu8r9nmnhns5gpdrgfjy" {
        t.Fatalf("Record is wrong")
    }
}

func TestResolveStandardUpper(t *testing.T) {
    client := buildClient()
    record, success := client.ResolveStandard("AVVY-CLIENT-COMMON-TESTING.AVAX", client.RECORDS["X_CHAIN"])
    if !success {
        t.Fatalf("Should succeed")
    }
    if record != "x-avax13fd740ykwc5peewmkcgu8r9nmnhns5gpdrgfjy" {
        t.Fatalf("Record is wrong")
    }
}

func TestResolveStandardNoResolver(t *testing.T) {
    client := buildClient()
    _, success := client.ResolveStandard("avvy-client-common-no-resolver.avax", client.RECORDS["X_CHAIN"])
    if success {
        t.Fatal("TestResolveStandardNoResolver: Should fail")
    }
}

func TestGetDomainAndHash(t *testing.T) {
    client := buildClient()
    domain, hash := client.GetDomainAndHash("test.avvy-client-common-testing.avax")
    if hash.String() != "1523251939777174889952058621189833214337688882241909193899539985084513239604" {
        t.Fatal("TestGetDomainAndHash: Hash seems wrong")
    }
    if domain.String() != "5009886810970053750228119498024191690423831754312784118430229935127343039646" {
        t.Fatal("TestGetDomainAndHash: Domain seems wrong")
    }
}
