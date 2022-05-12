package services

import (
	"vibrato/model"
	"vibrato/repository"
)

var FavoriteService = newFavoriteService()

func newFavoriteService() *favoriteService {
	return &favoriteService{}
}

type favoriteService struct {
}

func (s favoriteService) AddFavorite(userId int64, videoId int64) error {
	err := repository.FavoriteRepository.AddFavorite(&model.FavoriteEntity{
		UserId:  userId,
		VideoId: videoId,
	})
	return err
}

func (s favoriteService) DeleteFavorite(userId int64, videoId int64) error {
	err := repository.FavoriteRepository.DeleteFavorite(&model.FavoriteEntity{
		UserId:  userId,
		VideoId: videoId,
	})
	return err
}

func (s favoriteService) GetFavoriteList(userId int64) ([]model.Video, error) {
	favoriteEntities, err := repository.FavoriteRepository.GetFavoriteList(userId)
	if err != nil {
		return nil, err
	}
	videos := make([]model.Video, 0)
	for _, favoriteEntity := range favoriteEntities {
		video := MapVideo(&favoriteEntity.Video)
		videos = append(videos, video)
	}
	return videos, nil
}

// IsFavorite 是否收藏
func (s favoriteService) IsFavorite(userId int64, videoId int64) bool {
	if userId <= 0 || videoId <= 0 {
		return false
	}
	return repository.FavoriteRepository.IsFavorite(userId, videoId)
}

func (s favoriteService) FillVideosFavoriteStatus(videos []model.Video, userId int64) {
	for index, _ := range videos {
		videoId := videos[index].Id
		videos[index].IsFavorite = s.IsFavorite(userId, videoId)
	}
}

func mapVideos(videoEntities []*model.VideoEntity) []model.Video {
	videos := []model.Video{}
	for _, v := range videoEntities {
		videos = append(videos, model.Video{
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
		})
	}
	return videos
}
