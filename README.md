
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
    value, _ := client.ResolveStandard("avvydomains.avax", client.RECORDS["EVM"])
    fmt.Println(value)
}
```

# Development

See [DEVELOPMENT.md](DEVELOPMENT.md)
