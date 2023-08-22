package usecase

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gushikem01/go-handson/internals/config"
	"github.com/gushikem01/go-handson/internals/domain/model"
	"github.com/gushikem01/go-handson/internals/domain/repository"
)

type UserUsecase interface {
	FindUserById(c *gin.Context, id string) (*model.User, error)
	CreateUser(c *gin.Context, email, name, password string) (*model.User, error)
	UpdateUserById(c *gin.Context, id, email, name, password string) (*model.User, error)
	DeleteUserById(c *gin.Context, id string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
	tx       config.Transaction
}

func NewUserUsecase(
	userRepo repository.UserRepository,
	tx config.Transaction,
) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
		tx:       tx,
	}
}

// FindUserById ユーザー取得
func (u *userUsecase) FindUserById(c *gin.Context, id string) (*model.User, error) {
	// convert id to uuid
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return u.userRepo.FindUserById(c, uuid)
}

// CreateUser ユーザー作成
func (u *userUsecase) CreateUser(c *gin.Context, email, name, password string) (*model.User, error) {
	uuid := uuid.New()
	var user = &model.User{
		Id:       uuid,
		Email:    email,
		Name:     name,
		Password: password,
	}

	var err error
	var us *model.User
	// トランザクション開始
	err = u.tx.RunInTx(c, func(ctx context.Context) error {
		us, err = u.userRepo.CreateUser(ctx, user)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return us, nil
}

// UpdateUserById ユーザー更新
func (u *userUsecase) UpdateUserById(c *gin.Context, id string, email, name, password string) (*model.User, error) {

	var err error
	var us *model.User
	// ユーザー検索
	us, err = u.FindUserById(c, id)
	if err != nil {
		return nil, err
	}

	// ユーザー更新
	us = &model.User{
		Id:       us.Id,
		Email:    email,
		Name:     name,
		Password: password,
	}

	// トランザクション開始
	err = u.tx.RunInTx(c, func(ctx context.Context) error {
		us, err = u.userRepo.UpdateUser(ctx, us)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return us, nil
}

// DeleteUserById ユーザー削除
func (u *userUsecase) DeleteUserById(c *gin.Context, id string) error {
	// convert id to uuid
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return u.userRepo.DeleteUser(c, uuid)
}
