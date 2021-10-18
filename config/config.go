package config

import (
	"log"

	"github.com/kagi-opendoor222/go_practice3/utils"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	Logfile   string
	Static    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.Logfile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("name").MustString("8080"),
		Logfile:   cfg.Section("web").Key("logfile").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
