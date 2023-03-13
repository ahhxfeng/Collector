package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	Redis_host             string `json:"redis_host"`
	Redis_port             int    `json:"redis_port"`
	SignPipePrivateKeyPath string `json:"SignPipePrivateKeyPath"`
	PublicAddress          string `json:"PublicAddress"`
	PrivateAddress         string `json:"PrivateAddress"`
}

var Conf Config

func (conf *Config) LoadFromFile(path string) error {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		log.Println("errors: decoder json file failed")
		return err
	}
	return nil

}

func (conf *Config) SaveToFile(path string) error {
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	encoder := json.NewEncoder(file)
	err := encoder.Encode(&conf)
	if err != nil {
		log.Fatalln("errors: encoder json file failed")
		return err
	}
	return nil

}
