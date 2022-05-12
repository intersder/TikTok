package model

import (
	"database/sql"
	"time"
)

var Models = []interface{}{
	&UserEntity{}, &VideoEntity{}, &CommentEntity{}, &FavoriteEntity{},
}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

type UserEntity struct {
	Model
	Username      sql.NullString `gorm:"size:32;unique;not null;" json:"username" form:"username"`        // 用户名
	Name          string         `gorm:"size:32;not null;" json:"name" form:"name"`                       // 昵称
	Password      string         `gorm:"size:512;not null;" json:"password" form:"password"`              // 密码
	FollowCount   int64          `gorm:"not null;default:0;" json:"follow_count" form:"follow_count"`     // 关注数量
	FollowerCount int64          `gorm:"not null;default:0;" json:"follower_count" form:"follower_count"` // 粉丝数量
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`                                    // 创建时间
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`                                    // 更新时间
}

type VideoEntity struct {
	Model
	UserId        int64 `json:"user_id" form:"user_id"`
	User          UserEntity
	PlayUrl       string    `gorm:"size:512;not null;" json:"play_url" form:"play_url"`              // 视频地址
	CoverUrl      string    `gorm:"size:512;not null;" json:"cover_url" form:"cover_url"`            // 视频封面
	FavoriteCount int64     `gorm:"not null;default:0;" json:"favorite_count" form:"favorite_count"` // 获赞数量
	CommentCount  int64     `gorm:"not null;default:0;" json:"comment_count" form:"comment_count"`   // 评论数量
	CreatedAt     time.Time `json:"created_at" form:"created_at"`                                    // 创建时间
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`                                    // 更新时间
}

type CommentEntity struct {
	Model
	UserId    int64 `json:"user_id" form:"user_id"`   // 评论用户id
	VideoId   int64 `json:"video_id" form:"video_id"` // 评论视频id
	User      UserEntity
	Content   string    `json:"content"`                      // 评论内容
	CreatedAt time.Time `json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"` // 更新时间
}

type FavoriteEntity struct {
	UserId    int64 `gorm:"primaryKey;autoIncrement:false" json:"user_id" form:"user_id"`   // 评论用户id
	VideoId   int64 `gorm:"primaryKey;autoIncrement:false" json:"video_id" form:"video_id"` // 评论视频id
	Video     VideoEntity
	CreatedAt time.Time `json:"created_at" form:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"` // 更新时间
}
