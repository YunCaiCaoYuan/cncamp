## service:对内访问
    在主机内访问:curl http://10.101.127.73/healthz，显示200

## ingress:对外访问
    腾讯云主机，需要将type=loadBalancer改为NodePort
    k edit svc/ingress-nginx-controller -n ingress

    在浏览器上访问:http://ingress.sunbin123.xyz:32200/healthz，显示200。
