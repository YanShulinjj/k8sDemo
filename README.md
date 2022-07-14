### k8s 实战部署

app 服务 需要： mysql + redis

- myappv1_1.yaml: 配置了mysql 和redis
- myappv1_2.yaml: 配置了app镜像
- myappv1_3.yaml: 配置了hpa控制器，实现自动扩缩容

坑：

1. 如果hpa 一直unkown，查看hpa是否正确配置，并且hpa create之前deployment一定是已经存在的。
2. mysql 挂载的nfs一定要存在，并且要空。
3. 如果集群网络出现问题，尝试重启虚拟机，（生产环境勿用）

Step:

1. 构建docker镜像

   ```bash
   docker build -t app:1.1 .
   ```

2. 依次创建k8s内部资源

   ```bash
   kubectl create -f myappv1_1.yaml
   kubectl create -f myappv1_2.yaml
   kubectl create -f myappv1_3.yaml
   # 查看app占用cpu
   kubectl get hpa -n go
   # 查看pod
   kubectl get pod -n go
   # 查看svc,
   kubectl get svc -n go
   ```

3. 测试服务：

   ```bash
   curl IP:port/mysql/new?name=suyame
   > {message: "successfully!"}
   curl IP:port/mysql/get?name=suyame
   > ***一个结构体
   curl IP:port/redis/get?name=suyame
   > ***一个结构体
   ```



#### /just for learning





