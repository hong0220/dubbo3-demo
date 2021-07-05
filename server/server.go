package main

import (
	"context"
	"fmt"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	_ "dubbo.apache.org/dubbo-go/v3/common/proxy/proxy_factory"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/filter/filter_impl"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
	_ "dubbo.apache.org/dubbo-go/v3/registry/zookeeper"
	tripleConstant "github.com/dubbogo/triple/pkg/common/constant"
)

import (
	"dubbo3-demo/protobuf"
)

func init() {
	setConfigByAPI()
}

func setConfigByAPI() {
	providerConfig := config.NewProviderConfig(
		config.WithProviderAppConfig(config.NewDefaultApplicationConfig()),
		config.WithProviderProtocol("tripleProtocolKey", "tri", "20000"),
		// 注册中心
		config.WithProviderRegistry("registryKey", config.NewDefaultRegistryConfig("zookeeper")),

		config.WithProviderServices("greeterImpl", config.NewServiceConfigByAPI(
			config.WithServiceRegistry("registryKey"),
			config.WithServiceProtocol("tripleProtocolKey"),
			config.WithServiceInterface("org.apache.dubbo.UserProvider"),
		)),
	)
	config.SetProviderConfig(*providerConfig)
}

func main() {
	config.SetProviderService(NewGreeterProvider())
	config.Load()
	select {}
}

/**
 * 对外接口
 */
type GreeterProvider struct {
	*protobuf.GreeterProviderBase
}

/**
 * 对外接口对象
 */
func NewGreeterProvider() *GreeterProvider {
	return &GreeterProvider{
		GreeterProviderBase: &protobuf.GreeterProviderBase{},
	}
}

/**
 * 对外接口的处理方法
 */
func (s *GreeterProvider) SayHelloStream(svr protobuf.Greeter_SayHelloStreamServer) error {
	c, err := svr.Recv()
	if err != nil {
		return err
	}
	logger.Infof("Dubbo-go3 GreeterProvider recv 1 user, name = %s\n", c.Name)
	c2, err := svr.Recv()
	if err != nil {
		return err
	}
	logger.Infof("Dubbo-go3 GreeterProvider recv 2 user, name = %s\n", c2.Name)
	c3, err := svr.Recv()
	if err != nil {
		return err
	}
	logger.Infof("Dubbo-go3 GreeterProvider recv 3 user, name = %s\n", c3.Name)

	svr.Send(&protobuf.User{
		Name: "hello " + c.Name,
		Age:  18,
		Id:   "123456789",
	})
	svr.Send(&protobuf.User{
		Name: "hello " + c2.Name,
		Age:  19,
		Id:   "123456789",
	})
	return nil
}

/**
 * 对外接口的处理方法
 */
func (s *GreeterProvider) SayHello(ctx context.Context, in *protobuf.HelloRequest) (*protobuf.User, error) {
	logger.Infof("Dubbo3 GreeterProvider get user name = %s\n" + in.Name)
	fmt.Println("get triple header tri-req-id = ", ctx.Value(tripleConstant.TripleCtxKey(tripleConstant.TripleRequestID)))
	fmt.Println("get triple header tri-service-version = ", ctx.Value(tripleConstant.TripleCtxKey(tripleConstant.TripleServiceVersion)))
	return &protobuf.User{Name: "Hello " + in.Name, Id: "12345", Age: 21}, nil
}

/**
 * 对外接口的处理方法
 */
func (g *GreeterProvider) Reference() string {
	return "greeterImpl"
}
