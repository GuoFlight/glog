package main

import (
	"fmt"
	"github.com/GuoFlight/glog"
	"os"
)

func main(){
	logger,err := glog.NewLogger("./logs","DEBUG",false,3)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	logger.Info("hello glog")
}