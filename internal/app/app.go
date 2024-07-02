package app

import (
	"fmt"
	"goEdu/internal/config"
	"goEdu/internal/crud"
	"goEdu/internal/repos"
	"goEdu/internal/route"
	"goEdu/internal/service"
	"goEdu/pkg/auth"
	"goEdu/pkg/hash"
	"goEdu/pkg/logging"
)

func Run() {
	logger := logging.GetLogger()
	logger.Info("Start application")

	conf := config.GetConfig()

	tokenManager, err := auth.NewManager("qwerty")
	hasher := hash.NewSHA1Hasher("qwerty")

	repositories := repos.NewRepositories(crud.ConnPool)
	allService := service.NewServices(service.Deps{
		Repos:        repositories,
		TokenManager: tokenManager,
		Hasher:       hasher,
	})

	r := route.NewHandler(allService, tokenManager)
	err = r.Init().Run(fmt.Sprintf(":%s", conf.Server.Port))
	if err != nil {
		fmt.Println(err)
		panic("start app failure")
	}
}
