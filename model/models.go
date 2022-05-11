package model

import (
	"database/sql"
	"time"
)

var Models = []interface{}{
	&User{},
}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

type User struct {
	Model
	Username      sql.NullString `gorm:"size:32;unique;not null;" json:"username" form:"username"`        // 用户名
	Name          string         `gorm:"size:32;not null;" json:"name" form:"name"`                       // 昵称
	Password      string         `gorm:"size:512;not null;" json:"password" form:"password"`              // 密码
	FollowCount   int            `gorm:"not null;default:0;" json:"follow_count" form:"follow_count"`     // 关注数量
	FollowerCount int            `gorm:"not null;default:0;" json:"follower_count" form:"follower_count"` // 粉丝数量
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`                                    // 创建时间
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`                                    // 更新时间
}
