package biz

import (
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/mock_db"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/dao"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/model"
)

// CreateUser 创建用户
// 根据用户 id 去数据库查找记录。若有记录，则创建用户失败；若没有记录，则创建用户
func CreateUser(db *mock_db.MockDB, user *model.User) error {
	_, err := dao.FindUserByID(db, user.ID)
	if err != nil {
		return err
	}
	return dao.SaveUser(db, user)
}
