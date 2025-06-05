package command

import (
	"log"
	"net/http"

	"gitcode.com/JerichoYu/marmot/global"
	"gitcode.com/JerichoYu/marmot/route/v1Route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type (
	WebServiceCommand struct{}
)

var (
	WebService WebServiceCommand
)

func (*WebServiceCommand) New() *WebServiceCommand { return &WebService }

func (*WebServiceCommand) Run() {
	engine := gin.Default()
	engine.Use(cors.Default())

	{
		engine.Any("/heath", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
		engine.StaticFS("/upload/rezip", http.Dir("runtime/upload/rezip"))

		v1Route.Rezip.New().Register(engine)
	}

	// 启动web-service服务
	if err := engine.Run(":" + cast.ToString(global.Config.System.Port)); err != nil {
		global.Logger.Error("启动web服务", zap.Error(err))
		log.Fatalf("【启动web服务错误】%s", err.Error())
	}
}
