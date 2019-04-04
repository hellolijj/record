# Proposal for GSoC 2019

## Kubernetes with hardware devices topology awareness at node level

## About Me

- Name: Junjun Li

- Email: junjunli666@gmail.com

- Phone: +86 17826839787

- Github: [hellolijj](https://github.com/hellolijj)

- Education: 

  - 2018.9 ~ present: Software Engineering Master Candidate in Zhejiang University, China.

- Related Experience:

  - Successfully completed the requirements needed to achieve [CKA Certification](https://raw.githubusercontent.com/hellolijj/record/master/blog/gsoc/Cka_Certificate_lijj.png)

  - Five mouths Edge Compute Platform Project development experience, as a software engineer intern in Hangzhou Harmonycloud Technology Ltd.

## About Project

- Project Name: Kubernetes with hardware devices topology awareness at node level

- Project Description: [Kubernetes with hardware devices topology awareness at node level](https://github.com/cncf/soc/blob/master/README.md#kubernetes-with-hardware-devices-topology-awareness-at-node-level)

## Proposal

### Background

- [Device Plugin](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/)

  - Kubernetes provides a [device plugin framework](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/resource-management/device-plugin.md) for vendors to advertise their resources to the kubelet without changing Kubernetes core code. Instead of writing custom Kubernetes code, vendors can implement a device plugin that can be deployed manually or as a DaemonSet. The targeted devices include GPUs, High-performance NICs, FPGAs, InfiniBand, and other similar computing resources that may require vendor specific initialization and setup.

- [Node Topology Manager](https://github.com/kubernetes/community/blob/cd961fad1e9d1c2c2ddf85e36a2f93c202c1da89/contributors/design-proposals/node/topology-manager.md)  

  - In order to extract the best performance, optimizations related to CPU isolation and memory and device locality are required. However, in Kubernetes, these optimizations are handled by a disjoint set of components.

  - [Node Topology Manager](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/topology-manager.md#overview) provides a mechanism to coordinate fine-grained hardware resource assignments for different components in Kubernetes.

### [goals](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/0035-20190130-topology-manager.md#goals)

- Arbitrate preferred socket affinity for containers based on input from CPU manager and Device Manager.
- Provide an internal interface and pattern to integrate additional topology-aware Kubelet components.

### [Non-Goal](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/0035-20190130-topology-manager.md#non-goals)

- *Inter-device connectivity:* Decide device assignments based on direct device interconnects. This issue can be separated from socket locality. Inter-device topology can be considered entirely within the scope of the Device Manager, after which it can emit possible socket affinities. The policy to reach that decision can start simple and iterate to include support for arbitrary inter-device graphs.
- *HugePages:* This proposal assumes that pre-allocated HugePages are spread among the available memory nodes in the system. We further assume the operating system provides best-effort local page allocation for containers (as long as sufficient HugePages are free on the local memory node.
- *CNI:* Changing the Container Networking Interface is out of scope for this proposal. However, this design should be extensible enough to accommodate network interface locality if the CNI adds support in the future. This limitation is potentially mitigated by the possibility to use the device plugin API as a stopgap solution for specialized networking requirements.

## Design

### Synopsis

todo

## Schedule

### Preparation (Prior - May 7)

During this period, I will keep familiarizing myself with the kubernetes source code and focusing on the issue about *hardware topology awareness at node level*  

### Community bonding (May 7 - 27)

Over this period, I am going to discuss with community regarding my proposal and investigate how the existing hardware topology management work.

#### Discussion (May 7 - 17)

Discuss with my mentor on the feasibility of my proposal and consider if there are some other potenial enhancements to my plan.

#### Investigation(May 18 - 26)

Analyze some existing  hardware topology management code and think over how to add some new features  into code base.

### Coding (May 28 - August 20)

My goal over this period of time is to realize the improvement of the hard ware topology management.

#### Week 1 (May 28 - June 4)

Deeply  understand the detail design of topology manager.

#### Week 2 ( June 5 - June 11)

Implement the pod admit handler interface.

#### Week 3 (June 12 - June 18)

Implement participates in Kubelet pod admission.

#### Week 4 (June 19 - June 25)

Implement the topology manager interface a feature gate:

As much Topology Manager functionality as possible is stubbed when the feature gate is disabled.

#### Week 5 (June 26 - July 1)

Implement the topology manager interface a feature gate:

As much Topology Manager functionality as possible is stubbed when the feature gate is disabled.

#### Week 6 (July 2 - June 8)

Writing unit tests and documentation about *topology manager interface a feature gate*.

#### Week 7 (July 9 - July 15)

Add`GetTopologyHints()`method to CPU Manager. 

#### Week 8 (July 16 - July 22)

Writing unit tests and documentation about *add`GetTopologyHints()`method to CPU Manager*.

#### Week 9 (July 23 - July 29)

Add`GetTopologyHints()`method to Device Manager

#### Week 10 (July 30 - August 5)

Add Socket ID to Device structure in the device plugin interface.

#### Week 11 (August 6 - August 12)

Writing unit tests and documentation about *Add`GetTopologyHints()`method to Device Manager*.

#### Week 12 (August 13 - August 19)

Submit code and evaluations

## Extra Information

### how can i do it?

I will be based in Zhejiang University during the summer And I will be available 35 - 40 hours per week. Since our summer vacation is from June to September, so I believe that I have enough time to complete the project.

### why i choose gsoc?

I am an Open source lover, and GSoC provides me a chance to make contributions to open source projects with mentorship from great developers all over the world. I believe it is amazing and I definitely can learn quite a few cutting-edge techniques from that.
