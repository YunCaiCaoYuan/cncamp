## 容器化httpserver

### 构建镜像
    docker build -t 769460962/myhttpserver:v2.1 -f  ../mod_3/Dockerfile .

### 推送仓库
    docker push 769460962/myhttpserver:v2.1

### 运行镜像
    docker run -itd -p 80:80 769460962/myhttpserver:v2.1

### 查看IP配置
    sudo nsenter -t `ps aux |grep mod_two |grep -v grep | awk '{print $2}'` -n ip a

### 保存镜像文件
    docker save -o myhttpserve.tar 769460962/myhttpserver:v2.1

### 清理镜像(清理残存的、临时的、没有被使用的镜像文件)
    docker image prune
    它支持的子命令有：
        -a, --all: 删除所有没有用的镜像，而不仅仅是临时文件；
        -f, --force：强制删除镜像文件，无需弹出提示确认；
