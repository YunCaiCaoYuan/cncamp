## 作业要求：编写 Kubernetes 部署脚本将 httpserver 部署到 Kubernetes 集群，以下是你可以思考的维度。
    优雅启动:
        实现readinessProbe

    优雅终止
        GO代码监听sigterm信号

    资源需求和 QoS 保证
        指定了request和limits
    
    探活
        和就绪一样

    TODO:
    日常运维需求，日志等级
    配置和代码分离
