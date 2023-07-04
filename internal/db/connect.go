package db

import (
	"database/sql"
	"fmt"

	"github.com/gh0st3e/RedLab_Interview/internal/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func ConnectPSQL(logger *logrus.Logger, cfg *config.PSQLDatabase) (*sql.DB, error) {
	connStr := cfg.Address
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("couldn't open PSQL: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("couldn't ping PSQl: %s", err)
	}

	logger.Info("Ping PSQL - OK!")

	return db, nil
}
