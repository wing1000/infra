package base

import (
	"github.com/kataras/iris"
	"github.com/tietang/go-eureka-client/eureka"
	"imooc.com/resk/infra"
	"time"
)

type EurekaStarter struct {
	infra.BaseStarter
	client *eureka.Client
}

func (e *EurekaStarter) Init(ctx infra.StarterContext) {
	e.client = eureka.NewClient(ctx.Props())
	rpcPort := ctx.Props().GetDefault("app.rpc.port", "18082")
	e.client.InstanceInfo.Metadata.Map["rpcPort"] = rpcPort
	e.client.Start()
}

func (e *EurekaStarter) Setup(ctx infra.StarterContext) {
	info := make(map[string]interface{})
	info["startTime"] = time.Now()
	info["appName"] = ctx.Props().GetDefault("app.name", "resk")
	Iris().Get("/info", func(context iris.Context) {
		context.JSON(info)
	})
	Iris().Get("/health", func(context iris.Context) {
		health := eureka.Health{
			Details: make(map[string]interface{}),
		}
		health.Status = eureka.StatusUp
		context.JSON(health)

	})
}
