package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zjmnssy/ezlog"
)

func main() {
	/*
		config := common.Config{
			LogFilePath:   "",
			LogLevels:     63,
			LogFilePrefix: "",
			Modules:       []string{common.ModulesAll},
			IsTimeSplit:   true,
			SplitPeriod:   10,
			IsSizeSplit:   true,
			SplitSize:     1000,
			IsClear:       true,
			SavePeriod:    1,
			UnifyTo:       common.UnifyTypeOfScreen,
		}

		ezlog.Init(config)
	*/

	ezlog.Debugf("my log=%d", 111)
	ezlog.Infof("my log=%d", 111)
	ezlog.Warnf("my log=%d", 111)
	ezlog.Errorf("my log=%d", 111)
	ezlog.Errorf("my log=%d")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			{
				fmt.Printf("security exit by %s signal.\n", s)
				time.Sleep(time.Millisecond * time.Duration(1500))
				os.Exit(0)
			}
		default:
			{
				fmt.Printf("unknown exit by %s signal.\n", s)
				time.Sleep(time.Millisecond * time.Duration(1500))
				os.Exit(0)
			}
		}
	}
}
