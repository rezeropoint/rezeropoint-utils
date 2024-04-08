package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig(fileAddress string) {
	// 初始化配置文件
	// 设定要读取的配置文件的路径
	viper.SetConfigFile("./configs/" + fileAddress)
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalln("读取配置文件出错", err)
	}

	// 配置文件找到并成功解析
	logrus.Infoln("载入配置文件成功")
	// 实时读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		logrus.Infoln("重载配置文件:", e.Name)
	})
	// 如果开启DuBug模式，则日志输出等级也设置为DebugLevel
	if viper.GetString("debug") == "true" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debugln("已启用Debug模式")
	}
}
