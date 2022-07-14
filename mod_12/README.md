### Istio安装
    curl -L https://istio.io/downloadIstio | sh -
    istioctl install --set profile=demo -y
### 准备
    命名空间Istio注入:
      kubectl label ns default istio-injection=enabled
    删除之前的pod，重建Pod:
      kubectl delete pod httpserver-8d6ff89df-9wtxd
      ...
### http方式
    建一个http的gateway:
      kubectl create -f httpgw.yaml	
    部署一个virtualServer:
      kubectl create -f httpvs.yaml
