set -o errexit
set -o nounset

echo "----------------------------------"
echo "Now install kubernetes in aliyun"
echo "Is it the master node ?"
echo "(Y/y) Y"
echo "(N/n) N"
echo "(0) exit"
echo "----------------------------------"
read input

case $input in
    Y | y )
    MASTER_NODE="true";;
    N | n)
    MASTER_NODE="false";;
    0)
    exit;;
esac
echo "MASTER_NODE:"$MASTER_NODE

uninstall_docker_centos() {
    yum remove docker \
        docker-client \
        docker-client-latest \
        docker-common \
        docker-latest \
        docker-latest-logrotate \
        docker-logrotate \
        docker-selinux \
        docker-engine-selinux \
        docker-engine
}

install_docker_centos() {
    yum install -y yum-utils \
        device-mapper-persistent-data \
        lvm2
    
    yum-config-manager \
        --add-repo \
        https://download.docker.com/linux/centos/docker-ce.repo
    yum -y update
    yum install -y docker-ce

	systemctl enable docker
	systemctl start docker
}

install_kube_centos() {

    
    cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF

    # 使用setenforce 0 的原因是在kubelet操作docker的时候存在一个权限的问题
    # Set SELinux in permissive mode (effectively disabling it)
    setenforce 0
    # sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config

    yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes 
    systemctl enable kubelet && systemctl start kubelet
}

k8s_config() {
    cat <<EOF >  /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOF
    sysctl --system

    echo 1 >> /proc/sys/net/ipv4/ip_forward

    # enable firewalld
    systemctl disable firewalld && systemctl stop firewalld
}


setup_master() {
	kubeadm init --pod-network-cidr=10.244.0.0/16 --image-repository registry.aliyuncs.com/google_containers --ignore-preflight-errors=all
    # Config default kubeconfig for kubectl
    mkdir -p "${HOME}/.kube"
    cat /etc/kubernetes/admin.conf > "${HOME}/.kube/config"
    chown "$(id -u):$(id -g)" "${HOME}/.kube/config"

    until kubectl get nodes &> /dev/null; do echo "Waiting kubernetes api server for a second..."; sleep 1; done
    # Enable master node scheduling

    kubectl taint nodes --all  node-role.kubernetes.io/master- 
    install_flannel
}

# Install flannel for the cluster
install_flannel(){
    kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/bc79dd1505b0c8681ece4de4c0d86c5cd2643275/Documentation/kube-flannel.yml
}

install_docker_centos
install_kube_centos
k8s_config

if $MASTER_NODE; then
    setup_master
fi
