package repository

import (
	"errors"
	"gorm.io/gorm"
	"vibrato/model"
	"vibrato/sqls"
)

var CommentRepository = newCommentRepository()

func newCommentRepository() *commentRepository {
	return &commentRepository{}
}

type commentRepository struct {
}

func (r *commentRepository) Create(comment *model.CommentEntity) error {

	err := sqls.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.VideoEntity{}).
			Where("id = ?", comment.VideoId).
			Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r commentRepository) ListAllByVideoId(videoId int64) (comments []*model.CommentEntity, err error) {
	err = sqls.DB().
		Preload("User").
		Where("video_id = ?", videoId).
		Order("created_at desc").
		Find(&comments).Error
	return
}

func (r *commentRepository) Delete(commentId int64, userId int64) error {

	comm := &model.CommentEntity{}
	err := sqls.DB().Take(comm, commentId).Error

	if err != nil {
		return errors.New("comment not found")
	}
	if comm.UserId != userId {
		return errors.New("you are not the owner")
	}

	err = sqls.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", commentId).
			Delete(model.CommentEntity{}).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.VideoEntity{}).
			Where("id = ?", comm.VideoId).
			Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
