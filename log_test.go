package g_log

import (
	"testing"
)

func Test_New(t *testing.T) {
	err := New("/tmp/g_log.log", DEBUG, OTBASE)
	if err != nil {
		t.Errorf("create log object failed!")
	}

}

func Test_Info(t *testing.T) {
	New("/tmp/g_log.log", INFO, OTBASE)
	Info("%s\n", "test info")
	Debug("%s\n", "test info")
	DaemonFlush()
}

func Test_Debug(t *testing.T) {
	New("/tmp/g_log.log", DEBUG, OTBASE)
	Debug("%s\n", "test debug")
	DaemonFlush()
}

func Test_Warn(t *testing.T) {
	New("/tmp/g_log.log", WARN, OTBASE)
	Warn("%s\n", "test warn")
	DaemonFlush()
}

func Test_Error(t *testing.T) {
	New("/tmp/g_log.log", ERROR, OTBASE)
	Error("%s\n", "test error")
	DaemonFlush()
}
