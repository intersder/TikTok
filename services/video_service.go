package services

import (
	"vibrato/model"
	"vibrato/repository"
)

var VideoService = newVideoService()

func newVideoService() *videoService {
	return &videoService{}
}

type videoService struct {
}

func (s videoService) Publish(video *model.Video, userId int64) error {
	err := repository.VideoRepository.Create(&model.VideoEntity{
		UserId:   userId,
		PlayUrl:  video.PlayUrl,
		CoverUrl: video.CoverUrl,
	})
	return err
}

func (s videoService) ListByUserId(userId int64) ([]model.Video, error) {
	videoEntities, err := repository.VideoRepository.ListByUserId(userId)
	if err != nil {
		return nil, err
	}

	videos := mapVideos(videoEntities)
	return videos, nil
}

func (s videoService) Feed(latestTime int64) ([]model.Video, error) {
	videoEntities, err := repository.VideoRepository.List(10, latestTime)
	if err != nil {
		return nil, err
	}

	videos := mapVideos(videoEntities)
	return videos, nil
}

func MapVideos(videoEntities []*model.VideoEntity) []model.Video {
	videos := []model.Video{}
	for _, v := range videoEntities {
		videos = append(videos, MapVideo(v))
	}
	return videos
}

func MapVideo(v *model.VideoEntity) model.Video {
	return model.Video{
		Id:            v.Id,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    false,
		Author: model.User{
			Id:            v.User.Id,
			Name:          v.User.Name,
			FollowCount:   v.User.FollowCount,
			FollowerCount: v.User.FollowerCount,
			IsFollow:      false,
		},
	}
}
