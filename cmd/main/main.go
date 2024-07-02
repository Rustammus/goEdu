package main

import (
	"fmt"
	"goEdu/internal/config"
	"goEdu/internal/route"
	"goEdu/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Start application")

	conf := config.GetConfig()
	r := route.SetupRouter(*conf)

	r.Run(fmt.Sprintf(":%s", conf.Server.Port))
}
