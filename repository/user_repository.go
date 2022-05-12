package repository

import (
	"vibrato/model"
	"vibrato/sqls"
)

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

type userRepository struct {
}

func (r *userRepository) Create(user *model.UserEntity) (err error) {
	err = sqls.DB().Create(user).Error
	return
}

func (r *userRepository) GetUserByUsername(username string) *model.UserEntity {
	ret := &model.UserEntity{}
	if err := sqls.DB().Take(ret, "username = ?", username).Error; err != nil {
		return nil
	}
	return ret
}
