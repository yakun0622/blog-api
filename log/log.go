package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)

	//if cfg.IsRel() {
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//}
	//if cfg.IsDev() {
	//	Logger.SetFormatter(&logrus.TextFormatter{
	//		DisableColors:   false,
	//		FullTimestamp:   true,
	//		TimestampFormat: "2006-01-02 15:04:05",
	//	})
	//}
	Logger.SetOutput(os.Stdout)
}
