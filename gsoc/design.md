
# GPU SHARE 对 节点拓扑的支持

我觉得首先要在 list and watch 的时候上传节点拓扑信息
如下:

```
type DeviceNode struct {
	devices []*pluginapi.Device
    topology map[[2]Device] topologyType
}

type topologyType string

const (
    topologyTypeX = "X"
    topologyTypeSYS = "SYS"
    topologyTypeNODE = "NODE"
    topologyTypePHB = "PHB"
    topologyTypePXB = "PXB"
    topologyTypePIX = "PIX"
    topologyTypeNV = "NV#"
	
)

```

使用的时候在原gpushare的基础上添加的字段上`aliyun.com/gpu-men` 、 `aliyun.com/gpu-count` 不添加新的字段。当使用多个gpu的时候、通过数组的形式传递每个gpu 的使用信息。
如：调度分配给某个任务2个gpu，分别使用 2g、3g显存。
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: tensorflow-mnist
  labels:
    app: tensorflow-mnist
spec:
  replicas: 1
  selector:
    matchLabels:
      app: tensorflow-mnist
  template:
    metadata:
      labels:
        app: tensorflow-mnist
    spec:
      containers:
      - name: tensorflow-mnist
        image: cheyang/distributed-tf:1.6.0-gpu
        imagePullPolicy: "IfNotPresent"
        resources:
          limits:
            - aliyun.com/gpu-count: 2
            - aliyun.com/gpu-count: 3
```