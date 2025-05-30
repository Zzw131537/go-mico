package wrappers

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type userWrapper struct {
	client.Client
}

// 熔断
func (wrapper *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,   // 熔断器请求阈值
		ErrorPercentThreshold:  50,   // 错误百分比,当错误超过百分比时进行降级处理，直至熔断器再次开启
		SleepWindow:            5000, // 过多长时间，熔断器检测是否开启， ms
	}
	hystrix.ConfigureCommand(cmdName, config)

	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

func NewUserWrapper(c client.Client) client.Client {

	return &userWrapper{c}
}
