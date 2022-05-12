package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vibrato/model"
	"vibrato/services"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {

	_ = c.Query("user_id")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	token := c.Query("token")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	actionType, err := strconv.Atoi(actionTypeStr)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	if actionType != 1 && actionType != 2 {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "action_type must be 1 or 2"})
	}

	currentUser, err := services.TokenService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	if actionType == 1 {
		err = services.FavoriteService.AddFavorite(currentUser.Id, videoId)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		return
	}
	if actionType == 2 {
		err = services.FavoriteService.DeleteFavorite(currentUser.Id, videoId)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		return
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	currentUser, err := services.TokenService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	favoriteList, err := services.FavoriteService.GetFavoriteList(currentUser.Id)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	for i, _ := range favoriteList {
		favoriteList[i].IsFavorite = true
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: favoriteList,
	})
}
