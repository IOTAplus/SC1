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
$ git checkout develop
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
````
$ ps aux | grep wasp | awk '{print $2}' | xargs kill -9
```


