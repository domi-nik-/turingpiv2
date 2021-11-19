# How to install your kubernets cluster on a turingpi v2 with talos

## Prerequisites

### talosctl

```zsh
curl -Lo /usr/local/bin/talosctl https://github.com/talos-systems/talos/releases/latest/download/talosctl-$(uname -s | tr "[:upper:]" "[:lower:]")-amd64
chmod +x /usr/local/bin/talosctl
```

### kubectl

#### Debian based

```zsh
sudo apt-get update && sudo apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubectl

```

#### Redhat based

```zsh
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF
yum install -y kubectl
```

See https://kubernetes.io/de/docs/tasks/tools/install-kubectl/

## My network setup at home

![Network Setup](/images/TuringPiv2_network.png "foo")

If you have a simpler network setup that’s totally fine! The three components DSL-Router, WLAN AP and Switch will be all combined in your DSL-Router so just use the network ports of that. 

## Setup your board

Power supply, pins, cards, boards

## Prepare the OS Image

```zsh
curl -LO https://github.com/talos-systems/talos/releases/latest/download/metal-rpi_4-arm64.img.xz
xz -d metal-rpi_4-arm64.img.xz
sudo dd if=metal-rpi_4-arm64.img of=/dev/mmcblk0 conv=fsync bs=4M
```

As an alternativer you could you the raspberry pi imager tool and select the downloaded os image as input source.

## Check zour CM4s and SD-Cards

Test all sd cards and rpi on node port 1 of the tpi to check if all works fine.

If that looks ok to you, go and connect the turingpi v2 to your switch / dsl router with an LAN cable.


## Prepare your node configs

talosctl gen config "cluster-name" "cluster-endpoint"

### Adjustment the configs for you CM4s

Search and replace /dev/sda with /dev/mmcblk0 in the controplane.yaml and worker.yaml because the RaspberryPis don’t have an /dev/sda.

### What are the differences of the two configs?

"Control plane nodes in Talos run etcd which provides data store for Kubernetes and Kubernetes control plane components (kube-apiserver, kube-controller-manager and kube-scheduler).

Control plane nodes are tainted by default to prevent workloads from being scheduled to control plane nodes." (see https://www.talos.dev/docs/v0.13/guides/troubleshooting-control-plane/)

## Setup your nodes

### Find the IPs of your nodes

How to evaluate the IP of your newly booted rpi:

Check your network management tool like unifi dashboard
try a network scan if your network is 192.168.7.1/24 (have a look in your router, if needed)

```zsh
nmap -sP -oG - 192.168.7.1/24
```

In my case the searched IP of my first node is: 192.168.7.211 and 192.168.7.248 (as I have only two nodes active right now.)


### Provision you nodes

```zsh
talosctl apply-config --insecure \
    --nodes 192.168.7.248 \
   --file controlplane-1.yaml
 
talosctl apply-config --insecure \
    --nodes 192.168.7.211 \
    --file worker-1.yaml
```

This can easily last a longer time!
In the meanwhile copy your talos config to the default place

### Copy the talos config

```zsh
cp talosconfig  ~/.talos/config
```

So rthis config will be used as default for all your talosctl comands.

Now set the right endpoints in your talos config

```zsh
talosctl  --context=talos01 config endpoint 192.168.7.248   
```

### Check your Setup

```zsh
talosctl dmesg -f -n 192.168.7.211 
```


## Kubernetes Bootstrap

So your talos nodes are setup, let's bootstrap kubernetes on your nodes.
You will see, that's pretty easy and forward.

```zsh
talosctl bootstap -n 192.168.7.248
```

That's all, pretty amazing!

## Connect to your newly created Kubernetes Cluster

### Download your kubeconfig

```zsh
talosctl kubeconfig -n 192.168.7.248
```

Try some simple commands if the connection works:

```zsh
kubectl get nodes
```
TODO image

```zsh
kubectl get pods -A
```
TODO image



There is a talos dashboard to see some pixels jumping around:

```zsh
talosctl -n 192.168.7.211,192.68.7.248 dashboard  
```
TODO image

