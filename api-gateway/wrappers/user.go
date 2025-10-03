package wrappers

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
)

type userWrapper struct {
	client.Client
}

// 定义用户熔断
func (wrapper *userWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()

	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  50,
		SleepWindow:            4000,
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
