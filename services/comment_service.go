package services

import (
	"vibrato/model"
	"vibrato/repository"
)

var CommonService = newCommonService()

func newCommonService() *commonService {
	return &commonService{}
}

type commonService struct {
}

func (s commonService) Create(content string, videoId int64, userId int64) error {

	err := repository.CommentRepository.Create(&model.CommentEntity{
		UserId:  userId,
		VideoId: videoId,
		Content: content,
	})
	return err
}

func (s commonService) ListAllByVideoId(videoId int64) ([]model.Comment, error) {
	commentEntities, err := repository.CommentRepository.ListAllByVideoId(videoId)
	if err != nil {
		return nil, err
	}

	commons := mapCommons(commentEntities)
	return commons, nil
}

func (s commonService) Delete(commentId int64, userId int64) error {
	err := repository.CommentRepository.Delete(commentId, userId)
	return err
}

func mapCommons(commonEntities []*model.CommentEntity) []model.Comment {
	commons := []model.Comment{}
	for _, v := range commonEntities {
		commons = append(commons, model.Comment{
			User: model.User{
				Id:            v.User.Id,
				Name:          v.User.Name,
				FollowCount:   v.User.FollowCount,
				FollowerCount: v.User.FollowerCount,
			},
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	return commons
}
