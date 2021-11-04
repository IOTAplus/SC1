# SC1
In this repo code and documentation for IOTA+ initial smart contract,
we will be using Public testnet already available to validate our smart contracts,
but we can also use a local and private instalation. 

The base functionality of this smart contract is to convert IOTA to IEXP tokens. 
but first the Tokens must be created

## Creating IEXP tokens
We can do this via a GUI wallet or using the cli-wallet available on goshimmer repo
and using (``LocalSetup/config.json``) on the cli-wallet installation path, install cli-wallet from goshimmer releases
f.e (https://github.com/iotaledger/goshimmer/releases/download/v0.7.7/goshimmer-0.7.7_Linux_x86_64.tar.gz)

* Init wallet
```
cli-wallet init
``` 
* Request funds
```
cli-wallet request-funds
```
* Create IEXP tokens f.e
```
cli-wallet create-asset -name IEXP -symbol IEXP -amount 1000
```
* Other useful comands
```
cli-wallet balance
cli-wallet asset-info -id Am5bmpzNFfvkho6XebKBiXMPzdzjqNrFdnamre13hu1Q
cli-wallet address -list
```

## CrowdSale smart contract
### Requirements 
* Install wasmtime and wasm-pack
* This contract will be used to convert IOTAs to IEXP tokens
co will be on (you can use either Go or Rust to work):
```
crowdsale
```
### Contract development
* We use schema tool for developing the SC
* Init
```
schema -init 
```
* Generate Go/Rust code
```
schema -rust -go 
```
* Change the schema.yaml file and run again:
```
schema -rust -go 
```
### Wasm (buid/deploy SC)
* Run ``wasm-pack`` build inside your rust/go files 
* Deploy the contract f.e
```
wasp-cli chain deploy-contract wasmtime examplesc "examplesc"  contracts/wasm/examplesc/pkg/example_bg.wasm
```

### wasp-cli
* Wasp-cli is a tool developed to interact with wasp (smart contract nodes), you can use ``LocalSetup/wasp-cli.json`` as config reference
```
wasp-cli request-funds
```
* Seeing balance
```
wasp-cli balance
```
* Allocating funds
```
wasp-cli chain deposit IOTA:10000
```

### User oriented documentation
* https://iotaplus.atlassian.net/wiki/spaces/IOTAPLUSGA/pages/44335105/DRAFT+Smart+contracts+Implementation+steps
### Reference architecture
* https://blog.iota.org/iota-smart-contracts-beta-release/ 
