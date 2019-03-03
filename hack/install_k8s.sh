set -o errexit
set -o nounset

echo "----------------------------------"
echo "Now install kubernetes version v1.12"
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


echo "----------------------------------"
echo "Can it access google images?"
echo "(Y/y) Y"
echo "(N/n) N"
echo "(0) exit"
echo "----------------------------------"
read input

case $input in
    Y | y )
    IS_ACCESS_GOOGLE="true";;
    N | n)
    IS_ACCESS_GOOGLE="false";;
    0)
    exit;;
esac
echo "IS_ACCESS_GOOGLE:"$IS_ACCESS_GOOGLE

MASTER_CIDR="10.244.0.0/16"

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

    if( $IS_ACCESS_GOOGLE); then

        cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
exclude=kube*
EOF

    else

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
    fi

    # Set SELinux in permissive mode (effectively disabling it)
    # setenforce 0
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

    # enable firewalld
    systemctl disable firewalld && systemctl stop firewalld
}

images_tag() {
    image_pull=(
        "junjunli/kube-proxy:v1.13.1"
        "junjunli/coredns:1.2.6"
        "junjunli/pause:3.1"

        "junjunli/kube-scheduler:v1.13.1"
        "junjunli/kube-apiserver:v1.13.1"
        "junjunli/kube-controller-manager:v1.13.1"
        "junjunli/etcd:3.2.24"
        "junjunli/flannel:v0.10.0-amd64"
    )
    image_tag=(
        "k8s.gcr.io/kube-proxy:v1.13.1"
        "k8s.gcr.io/coredns:1.2.6"
        "k8s.gcr.io/pause:3.1"

        "k8s.gcr.io/kube-scheduler:v1.13.1"
        "k8s.gcr.io/kube-apiserver:v1.13.1"
        "k8s.gcr.io/kube-controller-manager:v1.13.1"
        "k8s.gcr.io/etcd:3.2.24"
        "quay.io/coreos/flannel:v0.10.0-amd64"
    )
    
    if $MASTER_NODE; then
        end=${#image_pull[*]}
    else
        end=3
    fi

    for((i=0; i<$end; i++));
    do
       docker pull ${image_pull[$i]}
       docker tag ${image_pull[$i]} ${image_tag[$i]}
       docker rmi ${image_pull[$i]}
    done
}


setup_master() {

	kubeadm init --pod-network-cidr $MASTER_CIDR --ignore-preflight-errors=all

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

if( ! $IS_ACCESS_GOOGLE); then
   images_tag
fi

if $MASTER_NODE; then
    setup_master
fi