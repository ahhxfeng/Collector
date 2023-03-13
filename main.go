package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	ConfigFilePath      string
	PrintExampleConfig  bool
	UpdateExampleConfig bool
	PublicMux           *http.ServeMux
	PrivateMux          *http.ServeMux
)

func init_flag() {
	flag.StringVar(&ConfigFilePath, "c", "collector.json", "配置文件相对或绝对路径")
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
		log.Println("load file success ")
	}
	if UpdateExampleConfig {
		fmt.Println("try to back up config")
		configBackup := ConfigFilePath + ".bak"
		err := os.Rename(ConfigFilePath, configBackup)
		if err != nil {
			log.Fatalln("Back up config failed")
		} else {
			log.Println("old config has been reame to :", configBackup)
			log.Println("active config will update to :")
			log.Println(fmt.Println(Conf))
		}
		if err := Conf.SaveToFile(ConfigFilePath); err != nil {
			log.Println("update configure file failed")
		} else {
			log.Println("configure file update done")
		}
		return true
	}

	log.Println("current active configure:", Conf)
	return
}

func main() {
	init_flag()
	flag.Parse()
	if doConfig() {
		return
	}

}
