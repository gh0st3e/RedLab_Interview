package main

import (
	"os"

	"github.com/gh0st3e/RedLab_Interview/internal/config"
	"github.com/gh0st3e/RedLab_Interview/internal/db"

	"github.com/sirupsen/logrus"
)

const (
	FileStorage   = "files"
	DirPermission = 0777
)

func main() {
	logger := logrus.New()

	_, err := os.ReadDir(FileStorage)
	if err != nil {
		err := os.Mkdir(FileStorage, DirPermission)
		if err != nil {
			logger.Fatalf("Error while creating file storage: %s", err.Error())
		}
	}

	cfg, err := config.Init()
	if err != nil {
		logger.Fatalf("Error while load config: %s", err.Error())
		return
	}

	psql, err := db.ConnectPSQL(logger, &cfg.PSQLDatabase)
	if err != nil {
		logger.Fatalf("Error while conneting DB: %s", err.Error())
		return
	}

	_ = psql
}
