package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gh0st3e/RedLab_Interview/internal/entity"

	"github.com/lib/pq"
)

const (
	PgUniqueEntry = "23505"
)

type UserStore interface {
	NewUser(ctx context.Context, user entity.User) (int, error)
	RetrieveUser(ctx context.Context, login, password string) (entity.User, error)
}

func (s *Service) RetrieveUser(ctx context.Context, user entity.User) (*entity.User, error) {
	s.logger.Info("[RetrieveUser] started")

	user, err := s.userStore.RetrieveUser(ctx, user.Login, user.Password)
	if err != nil {
		s.logger.Errorf("[RetrieveUser] error in store: %s", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("incorrect login or password")
		}
		return nil, fmt.Errorf("error while process request\n%w", err)
	}

	s.logger.Info(user)
	s.logger.Info("[RetrieveUser] ended")

	return &user, nil
}

func checkUnique(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == PgUniqueEntry
}
