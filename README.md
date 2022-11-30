
# Installation

`go get github.com/avvydomains/golang-client`

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

## Installation

1. Clone this repo
2. Clone submodules

## Using an alternative client common setup

Set the `AVVY_CLIENT_COMMON` environment variable to reference the client common folder you wish to use.

## Building ABI

Run `bash scripts/generate_abi.sh` from the main folder. This requires Python & uses the client-common submodule.

