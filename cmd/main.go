package main

import (
	"github.com/gh0st3e/RedLab_Interview/internal/config"
	"github.com/gh0st3e/RedLab_Interview/internal/db"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logger := logrus.New()

	os.Mkdir("files", 0777)

	cfg, err := config.Init()
	if err != nil {
		logger.Fatalf("Error while load config: %s", err.Error())
		return
	}

	os.Chdir("files")

	psql, err := db.ConnectPSQL(logger, &cfg.PSQLDatabase)
	if err != nil {
		logger.Fatalf("Error while conneting DB: %s", err.Error())
		return
	}

	_ = psql
}
