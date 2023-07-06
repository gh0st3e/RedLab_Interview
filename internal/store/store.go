package store

import (
	"database/sql"
	"fmt"
	"github.com/gh0st3e/RedLab_Interview/internal/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	//PgUniqueEntryErrorCode postgres code for check unique violation
	PgUniqueEntryErrorCode = "23505"
)

type Store struct {
	db         *sql.DB
	ctxTimeout time.Duration
}

func NewStore(cfg config.PSQLDatabase) (*Store, error) {
	connStr := cfg.Address
	db, err := sql.Open(cfg.Driver, connStr)
	if err != nil {
		return nil, fmt.Errorf("couldn't open PSQL: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("couldn't ping PSQl: %s", err)
	}

	logrus.Info("Ping PSQL - OK!")

	return &Store{
		db:         db,
		ctxTimeout: time.Second * time.Duration(cfg.Timeout),
	}, nil
}

func isUniqueViolation(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == PgUniqueEntryErrorCode
}
