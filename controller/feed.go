package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"vibrato/model"
	"vibrato/services"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	latestTimeStr := c.Query("latest_time")
	token := c.Query("token")

	currentUser, _ := services.TokenService.GetUserByToken(token)

	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		latestTime = 0
	}
	videos, err := services.VideoService.Feed(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	if currentUser != nil {
		services.FavoriteService.FillVideosFavoriteStatus(videos, currentUser.Id)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
