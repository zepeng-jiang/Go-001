package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/DB"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/model"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/service"
	"net/http"
)

var db *DB.MockDB
var UserAlreadyExistsErr = sql.ErrNoRows
var ParamterErr = errors.New("user id must be greater than zero")

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// CreateUserResponse 创建用户响应
type CreateUserResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CreateUserHandler
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeRequest(r)
	if err != nil {
		fmt.Printf("[handler]: %+v \n", err)
	}
	srv := service.NewUserService()
	err = srv.CreateUserService(db, user)
	generateResponse(err, w)
}

// generateResponse 生成响应
func generateResponse(err error, w http.ResponseWriter) {
	if err != nil && errors.Cause(err) == UserAlreadyExistsErr {
		fmt.Printf("[handler] user error: %+v \n", err)
		encodeResponse(9999, "user already exists, can not create", w)
		return
	}
	if err != nil && errors.Cause(err).Error() == ParamterErr.Error() {
		encodeResponse(400, errors.Cause(err).Error(), w)
		return
	}
	if err != nil && err != UserAlreadyExistsErr {
		fmt.Printf("[handler] user error: %+v \n", err)
		encodeResponse(500, "Internal Server Error", w)
		return
	}
	encodeResponse(200, "create user success", w)
}

// decodeRequest 解码请求
func decodeRequest(r *http.Request) (*model.User, error) {
	req := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, errors.Wrap(err, "[handler] decode create request failed")
	}
	return &model.User{
		ID:       req.ID,
		Name:     req.Name,
		Password: req.Password,
	}, nil
}

// encodeResponse 编码响应
func encodeResponse(code int, msg string, w http.ResponseWriter) {
	resp := &CreateUserResponse{
		Code:    code,
		Message: msg,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Printf("[handler] encode create user response failed")
		return
	}
}
