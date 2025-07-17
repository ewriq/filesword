package utils

import (
	"github.com/go-ini/ini"
	"log"
)

type ConfigList struct {
	Port     			string
	Database 			string
	Username 			string
	Password 			string
}

var Config ConfigList

func LoadConfig(path string) ConfigList {
	cfg, err := ini.Load(path)
	if err != nil {
		log.Fatalf("Failed to load config file: %v", err)
	}

	
	Config := ConfigList{
		Port:     			cfg.Section("server").Key("port").MustString(":9000"),
		Database: 			cfg.Section("db").Key("dsn").String(),
		Username: 			cfg.Section("server").Key("username").String(),
		Password: 			cfg.Section("server").Key("password").String(),
	}
	

	return Config
}
