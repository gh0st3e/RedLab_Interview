package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gh0st3e/RedLab_Interview/internal/entity"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	FileStorage   = "files"
	DirPermission = 0777

	PgUniqueEntry = "23505"
)

type UserActions interface {
	NewUser(ctx context.Context, user entity.User) (int, error)
	RetrieveUser(ctx context.Context, login, password string) (entity.User, error)
}

type UserService struct {
	logger    *logrus.Logger
	userStore UserActions
}

func NewUserService(logger *logrus.Logger, userStore UserActions) *UserService {
	return &UserService{
		logger:    logger,
		userStore: userStore,
	}
}

func (u *UserService) NewUser(ctx context.Context, user entity.User) error {
	u.logger.Info("[NewUser] started")

	userID, err := u.userStore.NewUser(ctx, user)
	if err != nil {
		u.logger.Errorf("[NewUser] error in store: %s", err.Error())
		if checkUnique(err) {
			return fmt.Errorf("user with this login already exists")
		}
		return fmt.Errorf("error while process request, try later\n%w", err)
	}

	strID := strconv.Itoa(userID)

	err = os.Mkdir(filepath.Join(FileStorage, strID), DirPermission)
	if err != nil {
		u.logger.Errorf("[NewUser] error while creating user storage: %s", err.Error())
		return fmt.Errorf("error while process request, try later\n%w", err)
	}

	u.logger.Info("[NewUser] ended")

	return nil
}

func (u *UserService) RetrieveUser(ctx context.Context, user entity.User) (*entity.User, error) {
	u.logger.Info("[RetrieveUser] started")

	user, err := u.userStore.RetrieveUser(ctx, user.Login, user.Password)
	if err != nil {
		u.logger.Errorf("[RetrieveUser] error in store: %s", err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("incorrect login or password")
		}
		return nil, fmt.Errorf("error while process request\n%w", err)
	}

	u.logger.Info(user)
	u.logger.Info("[RetrieveUser] ended")

	return &user, nil
}

func checkUnique(err error) bool {
	pgErr, ok := err.(*pq.Error)
	return ok && pgErr.Code == PgUniqueEntry
}
