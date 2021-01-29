package main

import (
	"MyLogger/Logger"
	"time"
)

func main() {
	//log := Logger.NewLog("debug")
	log := Logger.NewFileLogger("info", "./", "test.log", 10*1024*1024)
	for {
		log.Debug("msg")
		log.Warning("msg")
		id := 10010
		name := "理想"
		log.Error("msg, id:%d, name:%s", id, name)
		time.Sleep(2 * time.Second)
	}
}
