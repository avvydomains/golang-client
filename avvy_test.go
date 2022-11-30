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
