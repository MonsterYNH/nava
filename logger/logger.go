package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/MonsterYNH/nava/setting"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func InitLogger(config *setting.LoggerConfig) (*logrus.Logger, error) {
	logger := logrus.New()

	jsonFormatter := NewLoggerFormatter()

	logger.SetFormatter(jsonFormatter)

	switch strings.ToLower(config.LogLevel) {
	case "panic":
		logger.Level = logrus.PanicLevel
	case "fatal":
		logger.Level = logrus.FatalLevel
	case "error":
		logger.Level = logrus.ErrorLevel
	case "warn", "warning":
		logger.Level = logrus.WarnLevel
	case "info":
		logger.Level = logrus.InfoLevel
	case "debug":
		logger.Level = logrus.DebugLevel
	case "trace":
		logger.Level = logrus.TraceLevel
	default:
		return nil, fmt.Errorf("log level %s not vaild", config.LogLevel)
	}

	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		config.LogFileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(config.LogFileName),
		// 设置最大保存时间
		rotatelogs.WithMaxAge(time.Duration(config.LogMaxAge)*time.Hour*24),
		// 设置日志切割时间间隔
		rotatelogs.WithRotationTime(time.Duration(config.LogRotationTime)*time.Hour*24),
	)
	if err != nil {
		return nil, err
	}

	logger.SetOutput(logWriter)

	return logger, nil
}
