package main

import (
	"fmt"
	"github.com/gh0st3e/RedLab_Interview/internal/clients/token_service"
	"github.com/gh0st3e/RedLab_Interview/internal/config"
	"github.com/gh0st3e/RedLab_Interview/internal/handler"
	"github.com/gh0st3e/RedLab_Interview/internal/jwt"
	"github.com/gh0st3e/RedLab_Interview/internal/service"
	"github.com/gh0st3e/RedLab_Interview/internal/store"
	"github.com/gin-gonic/gin"

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

	tokenService := token_service.NewTokenService(logger, cfg.TokenServiceConfig)
	err = tokenService.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	jwtService := jwt.NewJWTService(logger, tokenService)

	Handler := handler.NewHandler(logger, commonService, jwtService)

	server := gin.New()

	Handler.Mount(server)

	err = server.Run(":" + cfg.Server.Port)
	if err != nil {
		logger.Fatalf("Couldn't run server: %s", err)
	}
}
