echo "Installing WASP prerequisites"
sudo apt update && sudo apt upgrade
sudo apt -y install \
  build-essential \
  libgflags-dev \
  libsnappy-dev \
  zlib1g-dev \
  libbz2-dev \
  liblz4-dev \
  libzstd-dev
git clone https://github.com/facebook/rocksdb.git
cd rocksdb || exit
git checkout v6.25.3
export CXXFLAGS='-Wno-error=deprecated-copy -Wno-error=pessimizing-move -Wno-error=class-memaccess'
make shared_lib
sudo make install-shared INSTALL_PATH=/usr
cd ..

echo "Installing GO"
sudo apt install wget software-properties-common apt-transport-https -y
wget https://golang.org/dl/go1.17.linux-amd64.tar.gz
sudo tar -zxvf go1.17.linux-amd64.tar.gz -C /usr/local/
echo "export PATH=/usr/local/go/bin:${PATH}" | sudo tee /etc/profile.d/go.sh
source /etc/profile.d/go.sh
rm go1.17.linux-amd64.tar.gz

echo "Installing WASP"
mkdir -p wasp-nodes && cd wasp-nodes || exit
git clone https://github.com/iotaledger/wasp.git
cd wasp || exit 
make install
git clone https://github.com/IOTAplus/SC1.git
cp sc1/testnet/wasp/config.json wasp-nodes/
cd wasp-nodes || exit
nohup wasp > /dev/null 2>&1 &
