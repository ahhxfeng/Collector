package main

import (
	"io"
	"log"
	"os"
)

func InitLogger() *log.Logger {
	file, _ := os.OpenFile("./log/collector.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	w := io.MultiWriter(os.Stdout, file)
	logger := log.New(w, "Collector:", log.Ldate|log.Ltime|log.Lshortfile)
	return logger

}

var Logger = InitLogger()
