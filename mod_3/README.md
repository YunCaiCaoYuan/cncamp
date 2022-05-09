## 容器化httpserver

### 构建镜像
    docker build -t 769460962/myhttpserver:v2.1 -f  ../mod_3/Dockerfile .

### 推送仓库
    docker push 769460962/myhttpserver:v2.1

### 运行镜像
    docker run -itd -p 80:80 769460962/myhttpserver:v2.1

### 查看IP配置
    sudo nsenter -t `ps aux |grep mod_two |grep -v grep | awk '{print $2}'` -n ip a
