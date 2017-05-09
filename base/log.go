package base

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type (
	LogBase struct {
		Fio   *os.File
		Level int
		Body  []byte
		Flock sync.Mutex
	}
)

func NewLogBase(fio *os.File, level int) *LogBase {
	body := make([]byte, 0)
	return &LogBase{Fio: fio, Level: level, Body: body}
}

func (this *LogBase) Debug(format string, v ...interface{}) {
	if this.Level == 1 {
		s := fmt.Sprintf(format, v...)
		s = timeNowStr() + "[DEBUG] " + s
		this.Flock.Lock()
		this.Body = append(this.Body, s...)
		this.Flock.Unlock()
		fmt.Println("debug")
	} else {
		panic("debug error!")
	}
}

func (this *LogBase) Info(format string, v ...interface{}) {
	if this.Level <= 2 {
		s := fmt.Sprintf(format, v...)
		s = timeNowStr() + "[INFO] " + s
		this.Flock.Lock()
		this.Body = append(this.Body, s...)
		this.Flock.Unlock()
	}
}

func (this *LogBase) Warn(format string, v ...interface{}) {
	if this.Level <= 3 {
		s := fmt.Sprintf(format, v...)
		s = timeNowStr() + "[WARN] " + s
		this.Flock.Lock()
		this.Body = append(this.Body, s...)
		this.Flock.Unlock()
	}
}

func (this *LogBase) Error(format string, v ...interface{}) {
	if this.Level <= 4 {
		s := fmt.Sprintf(format, v...)
		s = timeNowStr() + "[ERROR] " + s
		this.Flock.Lock()
		this.Body = append(this.Body, s...)
		this.Flock.Unlock()
	}
}

// 处理日志数据写入文件
// 文件配额
func (this *LogBase) Output() {
	this.Flock.Lock()
	_, err := this.Fio.Write(this.Body)
	if err != nil {
		log.Println("write file failed!", err.Error())
	}
	this.Fio.Sync()
	this.Body = nil
	this.Flock.Unlock()
}

func (this *LogBase) Close() error {
	return this.Fio.Close()
}

func timeNowStr() string {
	now := time.Now()
	format := "2006-01-02 15:04:05 Mon"
	return fmt.Sprint(now.Format(format))
}
