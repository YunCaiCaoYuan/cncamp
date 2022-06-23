## service:对内访问
    

## ingress:对外访问
    腾讯云主机，需要将type=loadBalancer改为NodePort
    k edit svc/ingress-nginx-controller -n ingress

    在浏览器上访问:http://ingress.sunbin123.xyz:32200/healthz，显示200。
