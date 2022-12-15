
# Installation

`go get github.com/avvydomains/golang-client/avvy`

# Usage

Forward Resolution

```go
import (
    "fmt"
    "github.com/avvydomains/golang-client/avvy"
)

func main() {
    rpcUrl := "https://api.avax.network/ext/bc/C/rpc"
    chainId := 43114
    client := new(avvy.Client)
    client.Init(rpcUrl, chainId)
    value, success := client.ResolveStandard("avvydomains.avax", client.RECORDS["EVM"])
    if success {
        fmt.Println(value)
    }
}
```

Reverse Resolution

```go
import (
    "fmt"
    "github.com/avvydomains/golang-client/avvy"
)

func main() {
    rpcUrl := "https://api.avax.network/ext/bc/C/rpc"
    chainId := 43114
    client := new(avvy.Client)
    client.Init(rpcUrl, chainId)
    hash, success := client.ReverseResolve(client.RECORDS["EVM"], "0x9BC4e7C1Fa4Ca66f6B2F4B6F446Dad80Ec541983")
    if !success {
        fmt.Println("Failed to reverse resolve value")
    }
    name, success := client.LookupHash(hash)
    if success {
        fmt.Println(name)
    } else {
        fmt.Println("Failed to look up hash")
    }
}
```

# Development

See [DEVELOPMENT.md](DEVELOPMENT.md)
