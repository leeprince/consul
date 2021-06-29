# 使用 dockerfile 通过 docker build 构建镜像
```
docker build -t consul-single .
```

# 构建 consul 容器，并运行 consul agent 的 server 模式 
```
consul agent  -server  -bootstrap-expect=1 -data-dir=/tmp/consul -node=dockerfileNode -bind=0.0.0.0  -ui  -client=0.0.0.0
```
