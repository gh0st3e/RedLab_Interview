package main

import (
	"github.com/gh0st3e/RedLab_Interview/internal/config"
	"github.com/gh0st3e/RedLab_Interview/internal/service"
	"github.com/gh0st3e/RedLab_Interview/internal/store"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	cfg, err := config.Init()
	if err != nil {
		logger.Fatalf("Error while load config: %s", err.Error())
		return
	}

	userStore, err := store.NewUserStore(cfg.PSQLDatabase)
	if err != nil {
		logger.Fatalf("Error while init user store: %s", err)
	}
	productStore, err := store.NewProductStore()
	if err != nil {
		logger.Fatalf("Error while init product store: %s", err)
	}

	commonService := service.NewService(logger, productStore, userStore)
	_ = commonService
}
