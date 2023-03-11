package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host       string `json: "host"`
	Port       int    `json: "port"`
	Redis_host string `json: "redis_host"`
	Redis_port int    `json: "redis_port"`
}

var Conf Config

func (Conf *Config) LoadFromFile(path string) error {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("errors: ")
		return err
	}
	fmt.Println("config", Conf)
	return nil

}
