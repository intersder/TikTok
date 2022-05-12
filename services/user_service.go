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
	userEntity := repository.UserRepository.GetUserByUsername(username)
	if userEntity == nil {
		return nil, errors.New("用户不存在")
	}
	if !passwd.Matches(userEntity.Password, password) {
		return nil, errors.New("密码错误")
	}
	return &model.User{
		Id:            userEntity.Id,
		Name:          userEntity.Name,
		FollowCount:   userEntity.FollowCount,
		FollowerCount: userEntity.FollowerCount,
		IsFollow:      false,
	}, nil
}

func (s *userService) Register(username, name, password string) (*model.User, error) {

	userEntity := &model.UserEntity{
		Username: sql.NullString{
			String: username,
			Valid:  len(username) > 0,
		},
		Name:     name,
		Password: passwd.EncodePassword(password),
	}

	err := repository.UserRepository.Create(userEntity)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Id:            userEntity.Id,
		Name:          userEntity.Name,
		FollowCount:   userEntity.FollowCount,
		FollowerCount: userEntity.FollowerCount,
		IsFollow:      false,
	}, nil
}
