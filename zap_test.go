package log

import (
	"github.com/cockroachdb/errors"
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	defer Sync()
	Debug("ddd")
	Info("iii")
	Warn("www")
	err := errors.New("test error")
	err2 := errors.Wrap(err, "wrap err test")
	//fmt.Printf("err9999: %+v\n", err2)
	Error("oe-p", zap.Error(err2), zap.String("instance_id", "xxx999"))
	//Error("oe", zap.Error(errors.Wrap(err, "wrap test")))

}
