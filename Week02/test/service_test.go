package test

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/DB"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/model"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/service"
	"testing"
)

var db *DB.MockDB
var UserAlreadyExistsErr = sql.ErrNoRows

func TestCreateUserServiceWithBadPath(t *testing.T) {
	user := &model.User{
		ID:       996,
		Name:     "打工仔",
		Password: "password",
	}
	srv := service.NewUserService()
	err := srv.CreateUserService(db, user)
	if err != nil && errors.Cause(err) == UserAlreadyExistsErr {
		fmt.Printf("[test] error: %+v \n", err)
		fmt.Println("----------分割线------------")
		t.Log("满足预期情况")
	} else {
		t.Error("不满足预期情况")
	}
}

func TestCreateUserServiceHappyPath(t *testing.T) {
	user := &model.User{
		ID:       965,
		Name:     "靓仔",
		Password: "password",
	}
	srv := service.NewUserService()
	err := srv.CreateUserService(db, user)
	if err == nil {
		t.Log("满足预期情况")
	} else {
		t.Error("不满足预期情况")
	}
}

func TestCreateUserServiceSadPath(t *testing.T) {
	except := "user id must be greater than zero"
	user := &model.User{
		ID:       0,
		Name:     "路人甲",
		Password: "password",
	}
	srv := service.NewUserService()
	err := srv.CreateUserService(db, user)
	actual := errors.Cause(err).Error()
	if err != nil && except == actual {
		fmt.Printf("[test] error: %+v \n", err)
		fmt.Println("----------分割线------------")
		t.Log("满足预期情况")
	} else {
		t.Error("不满足预期情况")
	}
}
