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
echo "node Cri will be 【"$Cri"】"

MASTER_CIDR="10.244.0.0/16"



start_master() {

    echo 1 >> /proc/sys/net/ipv4/ip_forward
    systemctl disable firewalld && systemctl stop firwalld
    setenforce 0
    swapoff -a

    
	if [ "$Cri" == "docker" ]; then
        kubeadm reset -f
        iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X
        kubeadm init --pod-network-cidr $MASTER_CIDR --ignore-preflight-errors=all
    else
        kubeadm reset -f --cri-socket=/var/run/pouchcri.sock
        iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X
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




