package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type ConfigObject struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbName"`
	} `yaml:"database"`
}

var Config = ConfigObject{}

func SetConfig() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
