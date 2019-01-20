function start_docker() {

    # kubeadm init
    kubeadm init --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=all

    # Config default kubeconfig for kubectl
    mkdir -p "${HOME}/.kube"
    cat /etc/kubernetes/admin.conf > "${HOME}/.kube/config"
    chown "$(id -u):$(id -g)" "${HOME}/.kube/config"

    # install flannel
    kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/bc79dd1505b0c8681ece4de4c0d86c5cd2643275/Documentation/kube-flannel.yml

    # install dashboard
    kubectl apply -f https://raw.githubusercontent.com/hellolijj/record/master/yaml/kubernetes-dashboard.yaml

    # print login token
    token=$(kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}') | grep 'token:' | awk '{print $2}')

    echo "login token: "$token
}

start_docker