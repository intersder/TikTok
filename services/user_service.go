package services

import (
	"database/sql"
	"errors"
	"vibrato/model"
	"vibrato/repository"
	"vibrato/utils/passwd"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (s userService) Login(username string, password string) (*model.User, error) {
	user := repository.UserRepository.GetUserByUsername(username)
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	if !passwd.Matches(user.Password, password) {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

func (s *userService) Register(username, name, password string) (*model.User, error) {

	user := &model.User{
		Username: sql.NullString{
			String: username,
			Valid:  len(username) > 0,
		},
		Name:     name,
		Password: password,
	}

	err := repository.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
