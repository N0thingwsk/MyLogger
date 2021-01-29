package Logger

import (
	"fmt"
	"time"
)

//日志对象
type ConsoleLogger struct {
	Level LogLevel
}

//构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := GetLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

//格式化并打印日志
func (c ConsoleLogger) log(format string, level LogLevel, arg ...interface{}) {
	if c.enable(level) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, line := getInfo(3)
		fmt.Printf("[%s]  [%s]  [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), LevelToString(level), fileName, funcName, line, msg)
	}
}

func (c ConsoleLogger) enable(level LogLevel) bool {
	return level >= c.Level
}

func (c ConsoleLogger) Debug(msg string, arg ...interface{}) {
	c.log(msg, DEBUG, arg...)
}

func (c ConsoleLogger) Info(msg string, arg ...interface{}) {
	c.log(msg, INFO, arg...)
}

func (c ConsoleLogger) Warning(msg string, arg ...interface{}) {
	c.log(msg, WARNING, arg...)
}

func (c ConsoleLogger) Error(msg string, arg ...interface{}) {
	c.log(msg, ERROR, arg...)
}

func (c ConsoleLogger) Fatal(msg string, arg ...interface{}) {
	c.log(msg, FATAL, arg...)
}
