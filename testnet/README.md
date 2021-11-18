# Requirements
* [golang](https://golang.org/doc/install)
* [wasmtime](https://wasmtime.dev)
* [tinygo](https://tinygo.org/getting-started/install/linux/#ubuntu-debian)

## Install dependencies
Install rockdb dependency as follows:
```
$ sudo apt update && sudo apt upgrade
$ sudo apt -y install \
  build-essential \
  libgflags-dev \
  libsnappy-dev \
  zlib1g-dev \
  libbz2-dev \
  liblz4-dev \
  libzstd-dev
$ git clone https://github.com/facebook/rocksdb.git
$ cd rocksdb
$ git checkout v6.25.3
$ export CXXFLAGS='-Wno-error=deprecated-copy -Wno-error=pessimizing-move -Wno-error=class-memaccess'
$ make shared_lib
$ sudo make install-shared INSTALL_PATH=/usr
```

# Public Goshimmer

## Creating IEXP tokens (cli-wallet)

### cli-wallet setup
In order to get the latest binary you can go to https://github.com/iotaledger/goshimmer/releases
and download `cli-wallet-0.8.0_Linux_x86_64.tar.gz` as an example and extract it to `~/bin`
```
$ wget https://github.com/iotaledger/goshimmer/releases/download/v0.8.0/cli-wallet-0.8.0_Linux_x86_64.tar.gz
$ tar -xvzf cli-wallet-0.8.0_Linux_x86_64.tar.gz
$ mkdir bin
$ cd bin
$ mv ../cli-wallet-0.8.0_Linux_x86_64/* .
$ chmod +x cli-wallet
```

Add the following to your `~/.bashrc`:
```
export PATH=$PATH:~/bin
```

Source the .bashrc file to make the changes permanent
```
$ source ~/.bashrc
```

Copy the file `sc1/testnet/goshimmer/config.json` to `~/bin/config.json`

### Initialize the wallet
```
$ cli-wallet init
```

This will create a new wallet and the state file `~/bin/wallet.dat`, in addition you will need to save your seed.

### Request funds
In order to get funds from the Goshimmer node that is setup in the `sc1/testnet/goshimmer/config.json` you need to run the following command:
```
$ cli-wallet request-funds
```

### Balance of the wallet
By default the balance is 1,000,000 IOTAs
```
$ cli-wallet balance
```

### Create 1000 IEXP assets
```
$ cli-wallet create-asset -name IEXP -symbol IEXP -amount 1000
```

### Get relevant information about the wallet
```
$ cli-wallet asset-info -id <asset-id>
$ cli-wallet address -list
```

# Private Wasp
Clone the repository with
```
$ git clone https://github.com/iotaledger/wasp.git
```

## wasp setup
Take into account the docs [here](https://wiki.iota.org/wasp/guide/chains_and_nodes/running-a-node)

```
$ cd wasp
$ make install
```

## Running the wasp node
```
$ cd
$ mkdir wasp-node
$ cp sc1/testnet/wasp/config.json wasp-node/
$ cd wasp-node
$ nohup wasp > /dev/null 2>&1 &
```
## How to stop the wasp node
```
$ ps aux | grep wasp | awk '{print $2}' | xargs kill -9
```

## Test the Dashboard
Go to the [dashboard](http://31.220.111.3:7000/) and login with `wasp/wasp`

## wasp-cli setup
```
$ cd
$ mkdir wasp-cli
$ cd wasp-cli
$ wasp-cli init
```

This will create a `wasp-cli.json` file in `~/wasp-cli/` that will only have the `wallet.seed` property in it. So you will need to add the `goshimmer` and `wasp` properties:
```
{
  "wallet": {
    "seed": "<my-seed>"
  },
  "goshimmer": {
    "api": "https://api.goshimmer.sc.iota.org"
  },
  "wasp": {
    "0": {
      "api": "127.0.0.1:9090",
      "nanomsg": "127.0.0.1:5550",
      "peering": "127.0.0.1:4000"
    }
  }
}
```

## Setting up a Chain (ISCP Chain)

### Trust setup
Here you can find the [docs](https://wiki.iota.org/wasp/guide/chains_and_nodes/setting-up-a-chain#trust-setup) to make the nodes trust to each other. Thi is tha case for more than one node operator. But we've seen that even with only a single node it's mandatory to also trust itself.

```
$ wasp-cli peering info
PubKey: JCwkc4WqKZzc3spRzWDA8jRFDF4YTo1rBezG8hBj6pge
NetID:  127.0.0.1:4000
$ wasp-cli peering trust JCwkc4WqKZzc3spRzWDA8jRFDF4YTo1rBezG8hBj6pge 127.0.0.1:4000
$ wasp-cli peering list-trusted
------                                        -----
PubKey                                        NetID
------                                        -----
JCwkc4WqKZzc3spRzWDA8jRFDF4YTo1rBezG8hBj6pge  127.0.0.1:4000
```

### Requesting test funds
```
$ wasp-cli request-funds --config wasp-cli/wasp-cli.json
Request funds for address 1BH6VTcoo2qND32oF1xaNqR5RKiPV4wh8xhP8WMqmG9hh: success
$ wasp-cli balance --config wasp-cli/wasp-cli.json 
Address index 0
  Address: 1CXdFSVdcLpeLyDvP3ZME9wYbXtJNxxmw7tFpdfxtCSvQ
  Balance:
    IOTA: 1000000
    ------
    Total: 1000000
```

### Deploy a chain
```
$ wasp-cli chain deploy \
  --committee=0 \
  --quorum=1 \
  --chain=iota-plus \
  --description="IOTA Plus Chain" \
  --config wasp-cli/wasp-cli.json
...
chain has been created successfully on the Tangle. ChainID: $/su8MqwXYTZkvbPtNZ34NvFQdQacaGronoJcC8WFdhpp5, State address: azJpRmAFgKgc2ZewLQgu5twWMPG5oBHMAhyf46EAaKbr, N = 1, T = 1
```

### Deposit funds to the Chain (10,000 IOTAs)
```
$ wasp-cli chain deposit \
  IOTA:10000 \
  --chain iota-plus \
  --config wasp-cli/wasp-cli.json
Posted on-ledger transaction F1LLGyctXvncauoZJHu3CRWxBxNfG4xhX12MfkgXctmX containing 1 request:
  - #0 (check result with: wasp-cli chain request 5hrXdCvzJXzmy9DQR3f1gDPrrAd5Auw7tYhPRabN6pB6kQw)
Waiting for tx requests to be processed...
```

## schema-tool setup
```
$ cd sc1
$ schema -init ERC20
$ cd erc20
$ go mod init github.com/iotaplus/SC1/erc20
$ go get github.com/iotaledger/wasp/packages/vm/wasmclient
$ go get github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib
$ schema -go
$ tinygo build -o erc20.wasm -target wasm go/main.go
```
## Deploy the smart contract
```
$ wasp-cli chain deploy-contract \
  wasmtime erc20 "ERC20 IEXP SC" \
  sc1/erc20/erc20.wasm \
  string o agentid A/1CXdFSVdcLpeLyDvP3ZME9wYbXtJNxxmw7tFpdfxtCSvQ::00000000 \
  string s int 10000000000 \
  --chain=iota-plus \
  --config wasp-cli/wasp-cli.json
uploaded blob to chain -- hash: F41ZuJTfpycQqHauVqVwQaPMLbJWCeHcmeHHisbumfpaPosted off-ledger request (check result with: wasp-cli chain request 5xcTGbnHcQKB1j4k6bx8pgh72AGg3e45guqcmmCT74kchps)
```

## EVM setup
### Setting up an EVM Chain
First of all you will need to deploy an [ISCP Chain](#deploy-a-chain) so the chain that we choose to create this time will be as follows:

```
$ wasp-cli chain deploy \
  --committee=0 \
  --quorum=1 \
  --chain=iscp-chain \    
  --description="ISCP Chain" \ 
  --config wasp-cli/wasp-cli.json
...
chain has been created successfully on the Tangle. ChainID: $/rF8hHRa1fELwTZu6nvt38nrsgRQgVxqkTMnDiYGVET2F, State address: L7wRpVhingmFbwQ64LLEc4bVnpkCxs8Cj6EWBswpHNQn, N = 1, T = 1
```

### Deposit funds to the new ISCP Chain (10,000 IOTAs by default)
```
$ wasp-cli chain deposit \
  IOTA:10000 \
  --config wasp-cli/wasp-cli.json \
  --chain iscp-chain
```

### Deploy the EVM Chain
```
$ wasp-cli chain evm deploy \
  --name evm-chain \
  --description "EVM Chain" \
  --alloc 0xc45AAA6AF36B81296d6Bbeb463d0130bC48e7567:1000000000000000000000000 \
  --config wasp-cli/wasp-cli.json \
  --chain iscp-chain
Posted off-ledger request (check result with: wasp-cli chain request 3saZUJ33ATRoSKdJLNbt8nroUS1PcoZpwNwMu7MhfJMrTBD)
evm-chain contract successfully deployed.
```

### Running the JSON-RPC Interface Server
```
$ nohup wasp-cli chain evm jsonrpc \
  --name evm-chain \
  --chainid 1074 \
  --config wasp-cli/wasp-cli.json > /dev/null 2>&1 &
```

### How to stop the JSON-RPC Interface Server
```
$ ps aux | grep wasp-cli | awk '{print $2}' | xargs kill -9
```