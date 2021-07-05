package main

import (
	"context"
	"time"
)

import (
	_ "dubbo.apache.org/dubbo-go/v3/cluster/cluster_impl"
	_ "dubbo.apache.org/dubbo-go/v3/cluster/loadbalance"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	_ "dubbo.apache.org/dubbo-go/v3/common/proxy/proxy_factory"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/filter/filter_impl"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
	_ "dubbo.apache.org/dubbo-go/v3/registry/zookeeper"
)

import (
	"dubbo3-demo/protobuf"
)

var greeterProvider = new(protobuf.GreeterClientImpl)

func init() {
	// 设置对外接口
	config.SetConsumerService(greeterProvider)
	setConfigByAPI()
}

func setConfigByAPI() {
	consumerConfig := config.NewConsumerConfig(
		config.WithConsumerAppConfig(config.NewDefaultApplicationConfig()),
		// 注册中心
		config.WithConsumerRegistryConfig("registryKey", config.NewDefaultRegistryConfig("zookeeper")),
		config.WithConsumerReferenceConfig("greeterImpl", config.NewReferenceConfigByAPI(
			config.WithReferenceRegistry("registryKey"),
			config.WithReferenceProtocol("tri"),
			config.WithReferenceInterface("org.apache.dubbo.UserProvider"),
		)),
	)
	config.SetConsumerConfig(*consumerConfig)
}

// need to setup environment variable "CONF_CONSUMER_FILE_PATH" to "conf/client.yml" before run
func main() {
	config.Load()

	time.Sleep(time.Second * 3)

	// 测试访问
	testSayHello()
}

func testSayHello() {
	ctx := context.Background()

	req := protobuf.HelloRequest{
		Name: "laurence",
	}
	user := protobuf.User{}
	err := greeterProvider.SayHello(ctx, &req, &user)
	if err != nil {
		panic(err)
	}

	logger.Infof("Receive user = %+v\n", user)
}
