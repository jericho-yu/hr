package v1Route

import (
	"github.com/gin-gonic/gin"

	"gitcode.com/JerichoYu/marmot/api/httpApi/v1HttpApi"
)

type (
	RezipRoute struct{}
)

var (
	Rezip RezipRoute
)

func (*RezipRoute) New() *RezipRoute { return &Rezip }

func (*RezipRoute) Register(engine *gin.Engine) {
	r := engine.Group("rezip")
	{
		r.POST("/upload", v1HttpApi.Rezip.Upload)
		
	}
}
