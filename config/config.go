package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	Logfile   string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("name").MustString("8080"),
		SQLDriver: cfg.Section("web").Key("logfile").String(),
		DbName:    cfg.Section("db").Key("driver").String(),
		Logfile:   cfg.Section("db").Key("name").String(),
	}
}
