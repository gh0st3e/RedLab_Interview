package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"time"

	"github.com/gh0st3e/RedLab_Interview/internal/config"
	"github.com/gh0st3e/RedLab_Interview/internal/entity"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	//PgUniqueEntryErrorCode postgres code for check unique violation
	PgUniqueEntryErrorCode = "23505"
)

type UserStore struct {
	db         *sql.DB
	ctxTimeout time.Duration
}

func NewUserStore(cfg config.PSQLDatabase) (*UserStore, error) {
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

	return &UserStore{
		db:         db,
		ctxTimeout: time.Second * time.Duration(cfg.Timeout),
	}, nil
}

// NewUser function which allows to create (register) new user (register)
func (s *UserStore) NewUser(ctx context.Context, user entity.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()

	query := `INSERT INTO users(login,password,name,email) VALUES($1,$2,$3,$4) RETURNING ID`

	var id int

	err := s.db.QueryRowContext(ctx, query,
		user.Login,
		user.Password,
		user.Name,
		user.Email).Scan(&id)

	if isUniqueViolation(err) {
		return 0, errors.New("user with this login already exists")
	}

	return id, err
}

// RetrieveUser func which allows to get user using login and password (login)
func (s *UserStore) RetrieveUser(ctx context.Context, login, password string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ctxTimeout)
	defer cancel()

	query := `SELECT id,name,email FROM users WHERE login=$1 AND password=$2`

	var user entity.User

	err := s.db.QueryRowContext(ctx, query, login, password).Scan(
		&user.ID,
		&user.Name,
		&user.Email)

	return user, err
}

func isUniqueViolation(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == PgUniqueEntryErrorCode
}
