package services

import (
	"context"
	"encoding/json"
	"errors"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
	"vibrato/model"
	"vibrato/sqls"
)

var TokenService = newTokenService()

func newTokenService() *tokenService {
	return &tokenService{}
}

type tokenService struct {
}

// GenerateToken 生成token
func (s *tokenService) GenerateToken(user *model.User) (string, error) {

	token := strings.ReplaceAll(uuid.NewV4().String(), "-", "")
	key := "login:token:" + token

	bytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	_, err = sqls.RDB().Set(context.Background(), key, string(bytes), time.Hour*24).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *tokenService) GetUserByToken(token string) (*model.User, error) {

	if len(token) == 0 {
		return nil, errors.New("token is empty")
	}

	key := "login:token:" + token
	var user model.User
	jsonStr, err := sqls.RDB().Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
