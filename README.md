# SC1
In this repo code and documentation for IOTA+ initial smart contract,
we will be using Public testnet already available to validate our smart contracts,
but we can also use a local and private instalation. 

The base functionality of this smart contract is to convert IOTA to IEXP tokens. 
but first the Tokens must be created

## Creating IEXP tokens
We can do this via a GUI wallet or using the cli-wallet available on goshimmer repo
and using (LocalSetup/config.json) on the cli-wallet installation path
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
* wasmtime and wasm-pack

* This contract will be used to convert IOTAs to IEXP tokens
```
CrowdSale.rst / CrowdSale.go
```
### wasp-cli
* wasp-cli is a tool developed to interact with wasp (smart contract nodes), you can use LocalSetup/wasp-cli.json as config reference
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
