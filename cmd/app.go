package main

import (
	"context"
	"dubbo-go-app/api"
	"dubbo-go-app/pkg/service"
	_ "dubbo-go-app/pkg/service"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"time"
)

// export DUBBO_GO_CONFIG_PATH=$PATH_TO_APP/conf/dubbogo.yaml
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	for {
		rsp, err := service.Client.SayHello(context.Background(), &api.HelloRequest{
			Name: "laurence",
		})

		logger.Warnf("GOT  RSP = %+v, err = %v\n", rsp, err)
		time.Sleep(time.Second)
	}
	select {

	}
}
