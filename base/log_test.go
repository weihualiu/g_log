package base

import (
	"errors"
	"os"
	"testing"
)

func Test_NewLogBase(t *testing.T) {
	fd, err := openFile()
	if err != nil {
		t.Errorf(err.Error())
	}
	lb := NewLogBase(fd, 1)
	if lb == nil || lb.Level != 1 {
		t.Errorf("NewLogBase init error!")
	}
	lb.Close()
}

func Test_Debug(t *testing.T) {
	fd, err := openFile()
	if err != nil {
		t.Errorf(err.Error())
	}
	lb := NewLogBase(fd, 1)
	if lb == nil || lb.Level != 1 {
		t.Errorf("NewLogBase init error for Debug!")
	}
	lb.Debug("%s\n", "test debug")
	if len(lb.Body) == 0 {
		t.Errorf("LogBase body set value failed!")
	}
	lb.Output()
	lb.Close()
}

func Test_Info(t *testing.T) {
	fd, err := openFile()
	if err != nil {
		t.Errorf(err.Error())
	}
	lb := NewLogBase(fd, 2)
	if lb == nil || lb.Level != 2 {
		t.Errorf("NewLogBase init error for Info!")
	}
	lb.Info("%s\n", "test info")
	if len(lb.Body) == 0 {
		t.Errorf("LogBase body set value failed!")
	}
	lb.Output()
	lb.Close()
}

func Test_Warn(t *testing.T) {
	fd, err := openFile()
	if err != nil {
		t.Errorf(err.Error())
	}
	lb := NewLogBase(fd, 3)
	if lb == nil || lb.Level != 3 {
		t.Errorf("NewLogBase init error for Warn!")
	}
	lb.Warn("%s\n", "test warn")
	if len(lb.Body) == 0 {
		t.Errorf("LogBase body set value failed!")
	}
	lb.Output()
	lb.Close()
}

func Test_Error(t *testing.T) {
	fd, err := openFile()
	if err != nil {
		t.Errorf(err.Error())
	}
	lb := NewLogBase(fd, 4)
	if lb == nil || lb.Level != 4 {
		t.Errorf("NewLogBase init error for error!")
	}
	lb.Error("%s\n", "test error")
	if len(lb.Body) == 0 {
		t.Errorf("LogBase body set value failed!")
	}
	lb.Output()
	lb.Close()
}

func openFile() (*os.File, error) {
	fd, err := os.OpenFile("/tmp/base_log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.New("NewLogBase open file error!")
	}
	return fd, nil
}
