package logger

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

type LoggerInfo struct {
	Time    string                 `json:"time"`
	Level   string                 `json:"level"`
	Message interface{}            `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type LogrusJsonFormatter struct{}

func NewLoggerFormatter() *LogrusJsonFormatter {
	return &LogrusJsonFormatter{}
}

func (formatter *LogrusJsonFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	loggerInfo := &LoggerInfo{
		Time:    entry.Time.Format("2006-01-02 15:04:05"),
		Level:   entry.Level.String(),
		Message: entry.Message,
		Data:    entry.Data,
	}

	bytesData, _ := json.Marshal(loggerInfo)

	return append(bytesData, '\n'), nil
}
