./build/setup.sh

sudo curl https://sh.rustup.rs -sSf | sh -s -- -y --default-toolchain=1.68.1
source $HOME/.cargo/env
rustup default 1.68.1

# Install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
k3d cluster create akri-cluster --servers 1