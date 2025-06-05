package global

import (
	"gitcode.com/JerichoYu/marmot/config"
	"go.uber.org/zap"
)

var (
	Config config.Config
	Logger *zap.Logger
)
