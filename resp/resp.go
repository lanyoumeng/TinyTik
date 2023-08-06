package resp

import (
	"github.com/gin-gonic/gin"
)

/*
200 OK: 请求成功，服务器成功处理了请求。
201 Created: 请求成功，并在服务器端创建了新的资源。
204 No Content: 请求成功，但服务器没有返回任何内容。
400 Bad Request: 请求有语法错误或参数错误，服务器无法处理该请求。
401 Unauthorized: 请求要求身份验证，用户没有提供有效的身份验证凭据，或者没有经过认证。
403 Forbidden: 服务器理解请求，但拒绝执行，用户没有权限访问特定资源。
404 Not Found: 请求的资源在服务器上不存在。
405 Method Not Allowed: 请求方法不允许被执行，通常用于限制特定请求方法的访问。
500 Internal Server Error: 服务器在执行请求时遇到了错误。
502 Bad Gateway: 服务器作为网关或代理，从上游服务器接收到无效的响应。
*/

// 基本回复
type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败 注：成功返回0
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

func Resp(c *gin.Context, httpCode int, data *Response) {
	c.JSON(httpCode, gin.H{
		"StatusCode": data.StatusCode,
		"StatusMsg":  data.StatusMsg,
	})
}