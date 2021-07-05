## 教程
https://dubbogo.github.io/zh-cn/docs/user/quickstart/3.0/quickstart.html

## 安装序列化工具protoc
https://github.com/protocolbuffers/protobuf/releases

## 安装 proto-gen-dubbo3 编译插件
```
export GO111MODULE="on"
export GOPROXY="http://goproxy.io"
go get -u github.com/apache/dubbo-go/protocol/dubbo3/protoc-gen-dubbo3@3.0
```

## 安装 dubbo3
```
export GOPROXY=https://mirrors.aliyun.com/goproxy/
go get dubbo.apache.org/dubbo-go/v3
```

## 使用安装好的编译工具编译接口
protoc-3.17.3-osx-x86_64/bin/protoc -I . protobuf/helloworld.proto --dubbo3_out=plugins=grpc+dubbo:.

## 启动zookeeper
docker-compose -f ./zookeeper.yml up -d

## 测试
分别启动服务端和客户端，可在客户端看到输出： 
```
2021-07-05T17:13:00.595+0800    INFO    client/client.go:64     Receive user = {Name:Hello laurence Id:12345 Age:21 XXX_NoUnkeyedLiteral:{} XXX_unrecognized:[] XXX_sizecache:0}
```
