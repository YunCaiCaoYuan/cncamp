## 修改记录:
    1、为 HTTPServer 添加 0-2 秒的随机延时
    已在mod_two添加；
    2、为 HTTPServer 项目添加延时 Metric；
    对应镜像：769460962/myhttpserver:v4.1
    3、将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
    已在mod_9的service.yaml metadata添加:
    annotations:
      prometheus.io/port: "80"
      prometheus.io/scrape: "true"
    部署Prometheus:
    helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    helm repo update
    helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack --create-namespace --namespace prometheus-stack
    4、从 Promethus 界面中查询延时指标数据；
    kubectl port-forward -n prometheus-stack  --address 0.0.0.0 svc/kube-prometheus-stack-prometheus 9090:9090
