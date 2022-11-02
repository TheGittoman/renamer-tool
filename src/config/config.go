package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config class for settings in json file
type Config struct {
	ScanFolder   string `json:"scanFolder"`
	FilterFolder string `json:"filterFolder"`
	SaveFolder   string `json:"saveFolderJson"`
}

func getConfig() *Config {
	conf := new(Config)
	file, err := ioutil.ReadFile("./data/config.json")
	if err != nil {
		file, err = ioutil.ReadFile("./data/config-default.json")
	}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// New make new Config object
func New() *Config {
	Config := getConfig()
	return Config
}

// func (c *Config) GetScanfolder() string {
// 	return c.ScanFolder
// }
// func (c *Config) GetFilterfolder() string {
// 	return c.FilterFolder
// }
// func (c *Config) GetSaveFolder() string {
// 	return c.SaveFolder
// }
