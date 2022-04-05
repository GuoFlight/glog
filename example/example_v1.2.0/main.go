package main

import (
	"fmt"
	"github.com/GuoFlight/glog"
	"os"
)

func main() {
	//logger, err := glog.NewLoggerWithConf(glog.DefaultConf)
	logger, err := glog.NewLoggerWithConf(glog.Conf{
		LogDir:            "./logs",
		LogLevel:          glog.InfoLevel,
		LogFileCount:      3,
		IsStdoutPrint:     false,
		CallerPrintLevels: nil,
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger.Info("hello glog")
}
