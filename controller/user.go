package controller

import (
	"TinyTik/common"
	"TinyTik/model"
	"TinyTik/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

// 用户信息 GET /douyin/user/
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	redis := common.GetRedisClient()
	if user, exist := redis.UserLoginInfo(token); exist {
		c.JSON(http.StatusOK, resp.UserResponse{
			Response: resp.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, resp.UserResponse{
			Response: resp.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
