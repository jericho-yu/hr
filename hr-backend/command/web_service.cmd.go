package command

import (
	"log"

	"gitcode.com/JerichoYu/marmot/global"
	"gitcode.com/JerichoYu/marmot/route"
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
	route.Index.NewInstance().Register(engine)

	// 启动web-service服务
	if err := engine.Run(":" + cast.ToString(global.Config.System.Port)); err != nil {
		global.Logger.Error("启动web服务", zap.Error(err))
		log.Fatalf("【启动web服务错误】%s", err.Error())
	}
}
