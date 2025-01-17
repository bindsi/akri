./build/setup.sh

sudo curl https://sh.rustup.rs -sSf | sh -s -- -y --default-toolchain=1.70
source $HOME/.cargo/env
rustup default 1.70

# Install k3d
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
k3d cluster create akri-cluster --servers 1
export AKRI_HELM_CRICTL_CONFIGURATION="--set kubernetesDistro=k3s"
helm repo add akri-helm-charts https://project-akri.github.io/akri/
helm install akri akri-helm-charts/akri $AKRI_HELM_CRICTL_CONFIGURATION

# Install go-lang
sudo apt-get update
sudo apt-get install -y golang-go