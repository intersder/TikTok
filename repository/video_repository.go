package repository

import (
	"time"
	"vibrato/model"
	"vibrato/sqls"
)

var VideoRepository = newVideoRepository()

func newVideoRepository() *videoRepository {
	return &videoRepository{}
}

type videoRepository struct {
}

// AddCommentCount 增加视频评论数
//func (r *videoRepository) AddCommentCount(videoId int64, num int) error {
//	err := sqls.DB().Model(&model.Video{}).
//		Where("id = ?", videoId).
//		Update("comment_count", gorm.Expr("comment_count + ?", num)).Error
//	return err
//}

func (r *videoRepository) Create(video *model.VideoEntity) (err error) {
	err = sqls.DB().Create(video).Error
	return
}

func (r videoRepository) ListByUserId(userId int64) (videos []*model.VideoEntity, err error) {
	err = sqls.DB().
		Preload("User").
		Where("user_id = ?", userId).
		Order("created_at desc").
		Find(&videos).Error
	return
}

func (r videoRepository) List(size int, latestTime int64) (videos []*model.VideoEntity, err error) {
	//err = sqls.DB().
	//	Preload("User").
	//	Limit(size).
	//	Where("created_at < ?", latestTime).
	//	Order("created_at desc").
	//	Find(&videos).Error

	db := sqls.DB().
		Preload("User").
		Limit(size)

	if latestTime > 0 {
		db.Where("created_at < ?", time.UnixMilli(latestTime))
	}
	err = db.Order("created_at desc").Find(&videos).Error
	return
}
