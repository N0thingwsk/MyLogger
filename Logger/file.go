package Logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//写日志文件

type FileLogger struct {
	Level       LogLevel
	filePath    string
	FileName    string
	errFileName string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := GetLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       LogLevel,
		filePath:    fp,
		FileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	file := path.Join(f.filePath, f.FileName)
	fileObj, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败", err)
		return err
	}
	errFileObj, err := os.OpenFile(file+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开错误日志文件失败", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(level LogLevel) bool {
	return level >= f.Level
}

func (f *FileLogger) log(format string, level LogLevel, arg ...interface{}) {
	if f.enable(level) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, line := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s]  [%s]  [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), LevelToString(level), fileName, funcName, line, msg)
		if level > ERROR {
			fmt.Fprintf(f.errFileObj, "[%s]  [%s]  [%s:%s:%d]  %s\n", now.Format("2006-01-02 15:04:05"), LevelToString(level), fileName, funcName, line, msg)
		}
	}
}

func (f *FileLogger) Debug(msg string, arg ...interface{}) {
	f.log(msg, DEBUG, arg...)
}

func (f *FileLogger) Info(msg string, arg ...interface{}) {
	f.log(msg, INFO, arg...)
}

func (f *FileLogger) Warning(msg string, arg ...interface{}) {
	f.log(msg, WARNING, arg...)
}

func (f *FileLogger) Error(msg string, arg ...interface{}) {
	f.log(msg, ERROR, arg...)
}

func (f *FileLogger) Fatal(msg string, arg ...interface{}) {
	f.log(msg, FATAL, arg...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
