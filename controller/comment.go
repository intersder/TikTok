package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vibrato/model"
	"vibrato/services"
)

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {

	_ = c.Query("user_id")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentIdStr := c.Query("comment_id")
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
		if len(commentText) == 0 {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "comment_text can't be empty"})
			return
		}
		err := services.CommonService.Create(commentText, videoId, currentUser.Id)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		return
	}
	if actionType == 2 {
		commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		err = services.CommonService.Delete(commentId, currentUser.Id)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
		return
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	//_ = c.Query("user_id")
	//token := c.Query("token")
	//_, _ = services.TokenService.GetUserByToken(token)

	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	comments, err := services.CommonService.ListAllByVideoId(videoId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: comments,
	})
}
