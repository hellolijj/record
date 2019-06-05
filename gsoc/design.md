
# GPU SHARE 对 节点拓扑的支持

我觉得首先要在 list and watch 的时候上传节点拓扑信息
如下:

```

type NvidiaDevicePlugin struct {
	devs         []*pluginapi.Device
	realDevNames []string
	devNameMap   map[string]uint
	devIndxMap   map[uint]string
	devTopology  map[string]map[string]topologyType
	
	socket       string
	mps          bool
	healthCheck  bool

	stop   chan struct{}
	health chan *pluginapi.Device

	server *grpc.Server
	sync.RWMutex
}

type topologyType nvml.P2PLinkType

```

使用的时候在原gpushare的基础上添加的字段上`aliyun.com/gpu-men` 、 `aliyun.com/gpu-count` 不添加新的字段。当使用多个gpu的时候、通过数组的形式传递每个gpu 的使用信息。
如：调度分配给某个任务需要使用2个gpu，分别使用 2g、3g显存。
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

> gpu share 实际上跟nvidia 的gpu 个数没什么两样。