# Requirements
* [golang](https://golang.org/doc/install)
* [wasmtime](https://wasmtime.dev/installation.html)
* [wasm-pack](https://rustwasm.github.io/wasm-pack/installer/)

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

Copy the file `testnet/goshimmer/config.json` to `~/bin/config.json`

### Initialize the wallet
```
$ cli-wallet init
```

This will create a new wallet and the state file `~/bin/wallet.dat`, in addition you will need to save your seed.

### Request funds
In order to get funds from the Goshimmer node that is setup in the `testnet/goshimmer/config.json` you need to run the following command:
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
$ wasp &
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

## Setting up a Chain

### Trust setup
Here you can find the [docs](https://wiki.iota.org/wasp/guide/chains_and_nodes/setting-up-a-chain#trust-setup) to make the nodes trust to each other. Thi is tha case for more than one node operator. But we've seen that even with only a single node it's mandatory to also trust itself.

```
$ wasp-cli peering info
PubKey: JCwkc4WqKZzc3spRzWDA8jRFDF4YTo1rBezG8hBj6pge
NetID:  0.0.0.0:4000
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
```

### Deploy a chain
```
$ wasp-cli chain deploy \
  --committee=0 \
  --quorum=0 \
  --chain=iexp-crowdsale \
  --description="IEXP Crowdsale" \
  --config wasp-cli/wasp-cli.json
...
chain has been created successfully on the Tangle. ChainID: $/su8MqwXYTZkvbPtNZ34NvFQdQacaGronoJcC8WFdhpp5, State address: azJpRmAFgKgc2ZewLQgu5twWMPG5oBHMAhyf46EAaKbr, N = 1, T = 1
```


