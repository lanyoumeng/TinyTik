package controller

import (
	"TinyTik/common"
	"TinyTik/resp"
	"TinyTik/service"
	"TinyTik/utils/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Res    resp.Response
	Videos []service.VideoList `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	title := c.PostForm("title")

	videoHeader, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Response{
			StatusCode: -1,
			StatusMsg:  "Get file err",
		})
		return
	}

	// 验证 token，获取 userID
	// userID, err := verifyToken(token)
	var userId int64
	token := c.PostForm("token")
	redis := common.GetRedisClient()
	if user, exist := redis.UserLoginInfo(token); exist {
		userId = user.Id
	} else {
		logger.Debug("user not exist")
	}

	err = service.NewVideo().Publish(c, title, videoHeader, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Response{
			StatusCode: -1,
			StatusMsg:  "Save file err",
		})

	} else {
		c.JSON(http.StatusOK, resp.Response{
			StatusCode: 0,
			StatusMsg:  videoHeader.Filename + " uploaded successfully",
		})
	}

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	videoService := service.NewVideo()
	videoList, err := videoService.PublishList(c, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, VideoListResponse{
			Res: resp.Response{
				StatusCode: -1,
				StatusMsg:  "Publish list false",
			},
			Videos: nil})

	} else {

		c.JSON(http.StatusOK, VideoListResponse{
			Res: resp.Response{
				StatusCode: 0,
				StatusMsg:  "Publish list success",
			},
			Videos: *videoList,
		})
	}
}
