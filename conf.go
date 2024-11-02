package glog

import "github.com/sirupsen/logrus"

type Conf struct {
	LogDir            string
	LogLevel          LogLevel
	LogFileCount      uint
	IsStdoutPrint     bool // 是否将日志输出到标准输出
	CallerPrintLevels []LogLevel
}

type LogLevel logrus.Level

const (
	PanicLevel LogLevel = LogLevel(logrus.PanicLevel)
	FatalLevel LogLevel = LogLevel(logrus.FatalLevel)
	ErrorLevel LogLevel = LogLevel(logrus.ErrorLevel)
	WarnLevel  LogLevel = LogLevel(logrus.WarnLevel)
	InfoLevel  LogLevel = LogLevel(logrus.InfoLevel)
	DebugLevel LogLevel = LogLevel(logrus.DebugLevel)
	TraceLevel LogLevel = LogLevel(logrus.TraceLevel)
)

var DefaultConf = Conf{
	LogDir:            "./logs",
	LogLevel:          InfoLevel,
	LogFileCount:      24,
	IsStdoutPrint:     false,
	CallerPrintLevels: []LogLevel{PanicLevel, FatalLevel, ErrorLevel, WarnLevel},
}
