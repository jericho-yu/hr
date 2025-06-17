package route

import (
	"net/http"

	"gitcode.com/JerichoYu/marmot/global"
	"gitcode.com/JerichoYu/marmot/route/v1Route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type (
	IndexRoute struct{}
)

var (
	Index IndexRoute
)

func (*IndexRoute) NewInstance() *IndexRoute { return &IndexRoute{} }

func (*IndexRoute) Register(engine *gin.Engine) {
	engine.Use(cors.Default())
	{
		engine.Any("/heath", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
		engine.StaticFS("/upload/rezip", http.Dir(global.Config.Rezip.OutDir))

		v1Route.Rezip.NewInstance().Register(engine)
	}
}
