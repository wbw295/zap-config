package log

import (
	zap_err "github.com/wbw295/zap-config/zap-err"
	"go.uber.org/zap"
	"testing"
)

func TestZapErr(t *testing.T) {
	defer Sync()
	err := zap_err.New("test err", zap.String("key1", "value1"))
	err = zap_err.Wrap(err, "vvv222", zap.String("key2", "value2"))
	err = zap_err.Wrap(err, "vvv333", zap.Uint("int", 9988))
	Error("field test6", zap_err.ToField(err))

}
