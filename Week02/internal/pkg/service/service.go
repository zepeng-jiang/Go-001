package service

import (
	"github.com/pkg/errors"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/DB"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/biz"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/model"
)

// Service
type Service interface {
	CreateUserService(db *DB.MockDB, user *model.User) error
}

type userService struct{}

// NewUserService 构造一个userService
func NewUserService() Service {
	return &userService{}
}

// CreateUserService 创建用户的service
func (us *userService) CreateUserService(db *DB.MockDB, user *model.User) error {
	if user.ID <= 0 {
		return errors.Wrap(errors.New("user id must be greater than zero"), "[service] request parameter wrong")
	}
	err := biz.CreateUser(db, user)
	if err != nil {
		return errors.WithMessage(err, "[service] create user failed")
	}
	return nil
}
