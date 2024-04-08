package log

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// MyFormatter 创建一个空的，供定义日志格式时使用
type MyFormatter struct{}

// Format 定义日志格式
func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	serverNmae := os.Getenv("POD_NAME")
	if serverNmae == "" {
		logrus.Fatalln("POD_NAME 环境变量未设置")
	}
	msg := fmt.Sprintf("[%s] [%s] [%s] %s: %s \n", serverNmae, strings.ToUpper(entry.Level.String()), timestamp, strings.ToUpper(entry.Caller.Function), entry.Message)
	return []byte(msg), nil
}

func InitLog() {
	// 初始化日志
	logrus.SetReportCaller(true)
	// 设置日志格式
	logrus.SetFormatter(new(MyFormatter))
	// 输出路径
	path := "./logs/"
	// 下面配置日志每隔 24 小时轮转一个新文件，保留最近 7 天的日志文件，多余的自动清理掉。
	logfile, _ := rotatelogs.New(
		path+"log_%Y%m%d.log.go",
		rotatelogs.WithLinkName(path+"log.log.go"),
		rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	logrus.SetOutput(io.MultiWriter(os.Stdout, logfile))
	logrus.SetLevel(logrus.InfoLevel)
}
