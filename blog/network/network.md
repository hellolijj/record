# 网络

## 容器网络

#### Network Namespace （网络栈）

- 默认情况下，我们在主机上安装某个软件“比如 apache” 会使用默认网络栈的 80 端口。如果再安装 apache 使用 80 端口，就会报错说端口被占用。
- 一个网络栈（namespace）包括 ： 网卡、回环设备、路由表和iptable规则
- 默认情况下，docker 每创建一个容器就会创建一个 network namespace

#### Bridge 网桥

- 主要根据 MAC 地址学习来将数据转发到网桥的不同端口
- docker 启动会检查主机是否有 docker0 网桥，如果没有则创建

#### Veth Pair

- Veth Pair 是虚拟网卡
- 成对出现，放入不同的Network Namespace 中

#### Overlay 网络

- 一种软件的形式，是整个集群的 “公有” 网络
- 一种形式的实现，可以通过 配置主机的路由表来实现 