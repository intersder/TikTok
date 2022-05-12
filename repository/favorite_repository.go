package repository

import (
	"gorm.io/gorm"
	"vibrato/model"
	"vibrato/sqls"
)

var FavoriteRepository = newFavoriteRepository()

func newFavoriteRepository() *favoriteRepository {
	return &favoriteRepository{}
}

type favoriteRepository struct {
}

// GetFavoriteList returns favorite list.
func (r *favoriteRepository) GetFavoriteList(userId int64) ([]*model.FavoriteEntity, error) {
	var favorites []*model.FavoriteEntity
	err := sqls.DB().Preload("Video").Preload("Video.User").
		Where("user_id = ?", userId).Find(&favorites).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *favoriteRepository) AddFavorite(favorite *model.FavoriteEntity) error {

	err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(favorite).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.VideoEntity{}).Where("id = ?", favorite.VideoId).
			Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *favoriteRepository) DeleteFavorite(favorite *model.FavoriteEntity) error {
	err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(favorite).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.VideoEntity{}).Where("id = ?", favorite.VideoId).
			Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r favoriteRepository) ListAllByVideoId(videoId int64) (favorites []*model.FavoriteEntity, err error) {
	err = sqls.DB().
		Preload("User").
		Where("video_id = ?", videoId).
		Order("created_at desc").
		Find(&favorites).Error
	return
}

// IsFavorite 查询用户是否收藏了该视频
func (r favoriteRepository) IsFavorite(userId int64, videoId int64) bool {
	count := int64(0)
	_ = sqls.DB().Model(&model.FavoriteEntity{}).
		Where("user_id = ?", userId).
		Where("video_id = ?", videoId).
		Count(&count).Error
	return count > 0
}
