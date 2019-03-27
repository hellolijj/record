#!/bin/bash
set -o errexit
set -o nounset

echo "----------------------------------"
echo "Is this a docker cri?"
echo "(Y/y) Y => docker"
echo "(N/n) N => pouch"
echo "(0) exit"
echo "----------------------------------"
read input

case $input in
    Y | y )
    Cri="docker";;
    N | n)
    Cri="pouch";;
    0)
    exit;;
esac
echo "change node 【"$node"】 Cri to be 【"$Cri"】"

MASTER_CIDR="10.244.0.0/16"

start_master() {
	if [ "$Cri" == "docker" ]; then
        kubeadm init --pod-network-cidr $MASTER_CIDR --ignore-preflight-errors=all
    else
        kubeadm init --pod-network-cidr $MASTER_CIDR --ignore-preflight-errors=all --cri-socket=/var/run/pouchcri.sock
    fi
    
    # Config default kubeconfig for kubectl
    mkdir -p "${HOME}/.kube"
    cat /etc/kubernetes/admin.conf > "${HOME}/.kube/config"
    chown "$(id -u):$(id -g)" "${HOME}/.kube/config"

    until kubectl get nodes &> /dev/null; do echo "Waiting kubernetes api server for a second..."; sleep 1; done
    # Enable master node scheduling

    kubectl taint nodes --all  node-role.kubernetes.io/master-
}

# Install flannel for the cluster
install_flannel(){
    kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/bc79dd1505b0c8681ece4de4c0d86c5cd2643275/Documentation/kube-flannel.yml
}

# install dashboard
install_dashboard() {
    kubectl apply -f https://raw.githubusercontent.com/hellolijj/record/master/yaml/dashboard/kubernetes-dashboard.yaml
}

start_master
install_flannel
install_dashboard