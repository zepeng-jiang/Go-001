package service

import (
	"github.com/pkg/errors"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/biz"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/mock_db"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/model"
)

// Service
type Service interface {
	CreateUserService(db *mock_db.MockDB, user *model.User) error
}

type userService struct{}

// NewUserService 构造一个userService
func NewUserService() Service {
	return &userService{}
}

// CreateUserService 创建用户的service
func (us *userService) CreateUserService(db *mock_db.MockDB, user *model.User) error {
	// 入参的校验应该放到之前，放到这里这里为了演示创建用户的 Sad Path
	if user.ID <= 0 {
		return errors.Wrap(errors.New("user id must be greater than zero"), "[service] request parameter wrong")
	}
	err := biz.CreateUser(db, user)
	if err != nil {
		return errors.WithMessage(err, "[service] create user failed")
	}
	return nil
}
