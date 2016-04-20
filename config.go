package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	Hostname       string `toml:"hostname"`
	Protocol       string `toml:"protocol"`
	APIPort        string `toml:"api_port"`
	FileServerPort string `toml:"fileserver_port"`
	WebRoot        string `toml:"webroot"`
	IndexTemplate  string `toml:"index_template"`
	IndexTarget    string `toml:"index_target"`
}

var config Settings

func loadConfig(configFile string) error {
	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		log.Fatal("Error loading config file")
		return err
	}
	return nil
}
