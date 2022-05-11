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

func (r *userRepository) Create(user *model.User) (err error) {
	err = sqls.DB().Create(user).Error
	return
}

func (r *userRepository) GetUserByUsername(username string) *model.User {
	ret := &model.User{}
	if err := sqls.DB().Take(ret, "username = ?", username).Error; err != nil {
		return nil
	}
	return ret
}
