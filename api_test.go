package ezlog_test

import (
	"testing"

	"github.com/zjmnssy/ezlog"
)

func TestDebugf(t *testing.T) {
	ezlog.Debugf("aaaa=%s", "aaa")
}
