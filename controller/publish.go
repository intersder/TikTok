package controller

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"path/filepath"
	"strings"
	"vibrato/model"
	"vibrato/services"
)

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if len(token) == 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "token不可为空"},
		})
		return
	}
	user, err := services.TokenService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	ext := filepath.Ext(strings.ToLower(data.Filename))
	allowExt := hashset.New(".mp4", ".mov", ".avi", ".flv",
		".wmv", ".rmvb", ".mkv", ".mts", ".m2ts", ".ts", ".m3u8")

	if !allowExt.Contains(ext) {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "文件格式不正确",
		})
		return
	}

	finalName := fmt.Sprintf("%d-%s%s", user.Id,
		strings.ReplaceAll(uuid.NewV4().String(), "-", ""), ext)

	// 视频文件上传
	url, err := services.UploadService.Upload(data, finalName)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	frameUrl, err := services.UploadService.GetSnapshot(finalName)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	err = services.VideoService.Publish(&model.Video{
		PlayUrl:  url,
		CoverUrl: frameUrl,
	}, user.Id)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {

	token := c.Query("token")
	//userIdStr := c.Query("user_id")
	//userId, err := strconv.ParseInt(userIdStr, 10, 64)
	//if err != nil {
	//	c.JSON(http.StatusOK, model.Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}

	if len(token) == 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "token不可为空"},
		})
		return
	}
	currentUser, err := services.TokenService.GetUserByToken(token)

	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	videos, err := services.VideoService.ListByUserId(currentUser.Id)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
