## service:对内访问
    在主机内访问:curl http://10.101.127.73/healthz，显示200

## ingress:对外访问
    腾讯云主机，需要将type=loadBalancer改为NodePort
    k edit svc/ingress-nginx-controller -n ingress

    在浏览器上访问:http://ingress.sunbin123.xyz:32200/healthz，显示200。

## Docker Pull设置代理解决Get https://k8s.gcr.io/v2/: net/http: request canceled while waiting for connection
    sudo mkdir -p /etc/systemd/system/docker.service.d
    sudo touch /etc/systemd/system/docker.service.d/proxy.conf
    sudo chmod 777 /etc/systemd/system/docker.service.d/proxy.conf
    sudo echo '
    [Service]
    Environment="HTTP_PROXY=http://proxy.xxx.com:8888/"
    Environment="HTTPS_PROXY=http://proxy.xxx.com:8888/"
    ' >> /etc/systemd/system/docker.service.d/proxy.conf
    sudo systemctl daemon-reload
    sudo systemctl restart docker

    通过systemctl查看配置结果：
    sudo systemctl show --property=Environment docker
