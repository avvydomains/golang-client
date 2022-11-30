
# Installation

`go get github.com/avvydomains/golang-client/avvy`

# Usage

```go
import (
    "github.com/avvydomains/golang-client/avvy"
)

func main() {
    rpcUrl := "https://api.avax.network/ext/bc/C/rpc"
    chainId := 43114
    client := new(avvy.Client)
    client.Init(rpcUrl, chainId)
    hash := client.NameHash("test.avax")
}
```

# Development

See [DEVELOPMENT.md](DEVELOPMENT.md)
