package log
import (
	"github.com/ermanimer/log/v2"
)
var Log *log.Logger

func InitLogs() {
  Log = log.NewLogger()
}
