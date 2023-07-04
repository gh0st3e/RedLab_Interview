package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/gh0st3e/RedLab_Interview/internal/entity"
)

const (
	CtxTimeout = 5 * time.Second
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

// NewUser function which allows to create (register) new user (register)
func (s *UserStore) NewUser(ctx context.Context, user entity.User) error {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeout)
	defer cancel()

	query := `INSERT INTO users(login,password,name,email) VALUES($1,$2,$3,$4)`

	_, err := s.db.ExecContext(ctx, query,
		user.Login,
		user.Password,
		user.Name,
		user.Email)

	return err
}

// RetrieveUser func which allows to get user using login and password (login)
func (s *UserStore) RetrieveUser(ctx context.Context, login, password string) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, CtxTimeout)
	defer cancel()

	query := `SELECT id,name,email FROM users WHERE login=$1 AND password=$2`

	var user entity.User

	err := s.db.QueryRowContext(ctx, query, login, password).Scan(
		&user.ID,
		&user.Name,
		&user.Email)

	return user, err
}
