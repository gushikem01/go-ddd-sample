package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/gushikem01/go-handson/internals/domain/model"
)

type UserRepository interface {
	FindUserById(c context.Context, id uuid.UUID) (*model.User, error)
	CreateUser(c context.Context, user *model.User) (*model.User, error)
	UpdateUser(c context.Context, u *model.User) (*model.User, error)
	DeleteUser(c context.Context, id uuid.UUID) error
}
