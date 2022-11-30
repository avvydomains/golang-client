## Installation

1. Clone this repo
2. Clone submodules

## Using an alternative client common setup

Set the `AVVY_CLIENT_COMMON` environment variable to reference the client common folder you wish to use.

## Building ABI

Run `bash scripts/generate_abi.sh` from the main folder. This requires Python & uses the client-common submodule to build .go files for each of the Avvy contracts.

## Generating contract addresses

Run `python scripts/generate_contract_addresses.py` to generate `avvy/contracts.json`. This is then bundled with the package for distribution.

If you are switching `AVVY_CLIENT_COMMON` you need to bundle again.
