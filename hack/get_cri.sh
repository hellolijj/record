    set -o errexit
    set -o nounset

    show_cri() {
        paste <(kubectl get node) <(cat <(echo "  Container Runtime Version") <(kubectl describe node | grep 'Container Runtime Version' | awk '{print "  "$4}'))
    }

    while(true); do show_cri; sleep 1; done

#paste <() <() 合并列
#cat <() <() 合并行
#awk '{print $a}'