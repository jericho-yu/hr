package main

import (
	"flag"
	"log"
	"strings"

	"gitcode.com/JerichoYu/marmot/command"
	"gitcode.com/JerichoYu/marmot/global"
	"github.com/fsnotify/fsnotify"
	"github.com/jericho-yu/aid/daemon"
	"github.com/jericho-yu/aid/zapProvider"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	var (
		originalCommand, commandName    string
		commandParams, originalCommands []string
		daemonOpen                      bool
	)
	flag.StringVar(&originalCommand, "t", "", "命令终端参数")
	flag.BoolVar(&daemonOpen, "daemon", false, "是否启动守护进程")
	flag.Parse()

	commandName = ""
	commandParams = make([]string, 0)

	if originalCommand != "" {
		originalCommands = strings.Split(originalCommand, " ")
		commandName = originalCommands[0]
		commandParams = originalCommands[1:]
	}

	launch(commandName, commandParams, originalCommands, daemonOpen)
}

// initLogger 初始化日志
func initLogger(logDir string) {
	var err error

	if global.Logger, err = zapProvider.ZapProviderApp.New(
		zapProvider.ZapProviderConfig.
			New(zap.ErrorLevel).
			SetEncoderType(zapProvider.EncoderTypeConsole).
			SetExtension(".log").
			SetInConsole(false).
			SetMaxSize(100).
			SetMaxDay(180).
			SetPathAbs(true).
			SetPath(logDir),
	); err != nil {
		log.Fatalf("【启动日志失败】 %s", err.Error())
	}
}

// initConfig 初始化配置
func initConfig(filename string) *viper.Viper {
	if filename == "" {
		filename = "./config.yaml"
	}

	v := viper.New()
	v.SetConfigFile(filename)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("【读取配置文件失败】%s", err.Error())
	}
	v.WatchConfig()

	if err = v.Unmarshal(&global.Config); err != nil {
		log.Fatalf("【初始化配置文件失败】%s", err.Error())
	}

	return v
}

// launch 启动程序
func launch(commandName string, commandParams, originalCommands []string, daemonOpen bool) {
	var (
		err error
		v   *viper.Viper
	)

	v = initConfig("")
	initLogger(global.Config.System.LogDir)
	v.OnConfigChange(func(in fsnotify.Event) {
		global.Logger.Info("配置文件改变", zap.String("file", in.Name))
		if err = v.Unmarshal(&global.Config); err != nil {
			global.Logger.Error("配置文件变更", zap.Error(err))
		}
	})

	if daemonOpen {
		daemon.App.New().Launch("启动程序", global.Config.System.LogDir) // 通过守护进程启动
	}

	switch commandName {
	default:
		command.WebService.New().Run()
	}
}
