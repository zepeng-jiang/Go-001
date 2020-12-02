package dao

import (
	"github.com/pkg/errors"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/mock_db"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/model"
)

// FindUserByID 通过 id 从数据库中查找用户
func FindUserByID(db *mock_db.MockDB, id int) (*model.User, error) {
	user, err := db.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "[dao] user not found ")
	}
	return user, nil
}

// SaveUser 将用户保存到数据库中
func SaveUser(db *mock_db.MockDB, user *model.User) error {
	err := db.Save(user)
	if err != nil {
		return errors.Wrap(err, "[dao] save user failed")
	}
	return nil
}
