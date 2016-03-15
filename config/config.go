package config

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

var All Settings

func Load(configFile string) {
	_, err := toml.DecodeFile(configFile, &All)
	if err != nil {
		log.Fatal("Error loading config file")
	}
}
