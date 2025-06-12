package v1HttpModule

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jericho-yu/aid/operation"
)

type (
	ResponseModule struct {
		C      *gin.Context
		Config ResponseConfig
	}

	ResponseConfig struct {
		HttpStatusCodeMode HttpStatusCodeModeType `json:"httpStatusCodeMode"`
	}

	HttpStatusCodeModeType = string

	SuccessResponseModule struct {
		Response *ResponseModule
		Code     int `json:"code"`
	}

	FailResponseModule struct {
		Response *ResponseModule
		Code     int `json:"code"`
	}
)

const (
	HttpStatusCodeModeFollowReal HttpStatusCodeModeType = "followReal"
)

var (
	Response        ResponseModule
	SuccessResponse SuccessResponseModule
	FailResponse    FailResponseModule
)

func (*ResponseModule) NewInstance(c *gin.Context, config ResponseConfig) *ResponseModule {
	return &ResponseModule{C: c, Config: config}
}

func (my *ResponseModule) Success() *SuccessResponseModule {
	return &SuccessResponseModule{Response: my, Code: 200}
}

func (my *ResponseModule) Fail() *FailResponseModule {
	return &FailResponseModule{Response: my, Code: 500}
}

func (my *SuccessResponseModule) Nil(msg string) {
	var message = "操作成功"

	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, my.Code),
		map[string]any{
			"code": my.Code,
			"msg":  message,
			"data": nil,
		},
	)
}

func (my *SuccessResponseModule) Data(msg string, data any) {
	var message = "success"
	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, my.Code),
		map[string]any{
			"code": my.Code,
			"msg":  message,
			"data": data,
		},
	)
}

func (my *SuccessResponseModule) Redirect(msg, url string) {
	var message = "redirect"

	if msg != "" {
		message = msg
	}

	my.Code = 301

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, my.Code),
		map[string]any{
			"code": my.Code,
			"msg":  message,
			"data": map[string]string{"url": url},
		},
	)
}

func (my *FailResponseModule) Nil(msg string) {
	var message = "操作失败"

	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, my.Code),
		map[string]any{
			"code": my.Code,
			"msg":  message,
			"data": nil,
		},
	)
}

func (my *FailResponseModule) UnAuth(msg string) {
	var message = "权限错误"

	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, 401),
		map[string]any{
			"code": 401,
			"msg":  message,
			"data": nil,
		},
	)
}

func (my *FailResponseModule) UnLogin(msg string) {
	var message = "未登录"

	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, 406),
		map[string]any{
			"code": 406,
			"msg":  message,
			"data": nil,
		},
	)
}

func (my *FailResponseModule) Empty(msg string) {
	var message = "不存在"

	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, 404),
		map[string]any{
			"code": 404,
			"msg":  message,
			"data": nil,
		},
	)
}

func (my *FailResponseModule) Forbidden(msg string) {
	var message = "禁止访问"

	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, 403),
		map[string]any{
			"code": 403,
			"msg":  message,
			"data": nil,
		},
	)
}

func (my *FailResponseModule) Error(msg string) {
	var message = "操作失败"

	if msg != "" {
		message = msg
	}

	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, 500),
		map[string]any{
			"code": 500,
			"msg":  message,
			"data": nil,
		},
	)
}

func (my *FailResponseModule) Format(format string, a ...any) {
	my.Response.C.JSON(
		operation.Ternary(my.Response.Config.HttpStatusCodeMode == "", 200, 500),
		map[string]any{
			"code": 500,
			"msg":  fmt.Sprintf(format, a...),
			"data": nil,
		},
	)
}
