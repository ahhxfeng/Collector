package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	ConfigFilePath      string
	PrintExampleConfig  bool
	UpdateExampleConfig bool
	PublicMux           *http.ServeMux
	PrivateMux          *http.ServeMux
)

func init_flag() {
	flag.StringVar(&ConfigFilePath, "c", "collector.config", "配置文件相对或绝对路径")
	flag.BoolVar(&PrintExampleConfig, "e", false, "打印一份样例配置，可以留作后用")
	flag.BoolVar(&UpdateExampleConfig, "u", false, "升级配置文件")
}

func doConfig() (quit bool) {
	if PrintExampleConfig {
		fmt.Println(Conf)
		return true
	}

	if err := Conf.LoadFromFile(ConfigFilePath); err != nil {
		log.Fatalln(err.Error())
	} else {

	}
	if UpdateExampleConfig {

	}
	return

}

func main() {
	init_flag()
	flag.Parse()
	if err := Conf.LoadFromFile(ConfigFilePath); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(Conf)
	}
}
