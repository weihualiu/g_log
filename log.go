package g_log

import (
	"errors"
	"github.com/weihualiu/g_log/base"
	"os"
	"time"
)

//
const (
	DEBUG = 1
	INFO
	WARN
	ERROR
)

const (
	OTBASE = 1
)

// 日志抽象接口
type Log interface {
	Debug(string, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
	Output()
	Close() error
}

// 全局对象
var logGlobal Log

// 以下是对外接口
func New(filepath string, level int, objtype int) error {
	fd, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("[g_log] read log file failed!")
	}
	if objtype == 1 {
		logGlobal = base.NewLogBase(fd, level)
	} else {
		return errors.New("[g_log] log object created failed!")
	}
	return nil
}

func DaemonFlush() {
	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			logGlobal.Output()
		}
	}()
	time.Sleep(time.Millisecond * 500)
}

func Infof(format string, v ...interface{}) {
	logGlobal.Info(format, v...)
}

func Debugf(format string, v ...interface{}) {
	logGlobal.Debug(format, v...)
}

func Warnf(format string, v ...interface{}) {
	logGlobal.Warn(format, v...)
}

func Errorf(format string, v ...interface{}) {
	logGlobal.Error(format, v...)
}

func Close() error {
	return logGlobal.Close()
}

func Info(v ...interface{}) {
	format := ""
	for range v {
		format += "%s"
	}
	format += "\n"
	Infof(format, v...)
}

func Debug(v ...interface{}) {
	format := ""
	for range v {
		format += "%s"
	}
	format += "\n"
	Debugf(format, v...)
}

func Warn(v ...interface{}) {
	format := ""
	for range v {
		format += "%s"
	}
	format += "\n"
	Warnf(format, v...)
}

func Error(v ...interface{}) {
	format := ""
	for range v {
		format += "%s"
	}
	format += "\n"
	Errorf(format, v...)
}
