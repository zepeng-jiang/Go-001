package mock_db

import (
	"database/sql"
	"github.com/zepeng-jiang/Go-000/Week02/pkg/model"
)

// MockDB mock数据库
type MockDB struct{}

// Save 保存
func (db *MockDB) Save(user *model.User) error {
	if nil != user {
		return nil
	}
	return sql.ErrNoRows
}

// FindByID 查询
func (db *MockDB) FindByID(id int) (*model.User, error) {
	if id == 996 {
		return nil, sql.ErrNoRows
	}
	return &model.User{
		Name:     "靓仔",
		Password: "password",
	}, nil
}
