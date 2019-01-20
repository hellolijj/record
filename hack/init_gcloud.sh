set -o errexit
set -o nounset

echo "----------------------------------"
echo "This is a script to init gcloud envirnments"
echo "it will install sshd docker and stop firewalld"
echo "----------------------------------"

esac

install sshd() {
    sudo -i
    # todo 
    vi /etc/ssh/sshd_config
    PermitRootLogin yes //默认为no，需要开启root用户访问改为yes

    # Change to no to disable tunnelled clear text passwords
    PasswordAuthentication yes //默认为no，改为yes开启密码登陆

    # todo repasswd

    systemctl restart sshd
}