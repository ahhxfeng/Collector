package common

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogger(name string) *logrus.Logger {
	// recive a file name than open for log, return a logrus.Logger
	file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	w := io.MultiWriter(file, os.Stdout)
	log := &logrus.Logger{
		Out: w,
		Formatter: &logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		},
		Level: logrus.TraceLevel,
	}

	return log

}

var Logger = InitLogger("./log/Collector.log")
