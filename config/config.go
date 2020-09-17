package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	configFile = "config.yaml"

	All    AllConfig
	Nasa   NASAConfig
	Server ServerConfig
)

type AllConfig struct {
	Nasa   NASAConfig   `yaml:"nasa"`
	Server ServerConfig `yaml:"server"`
}

type NASAConfig struct {
	APIURL  string `yaml:"api_url"`
	APIKey  string `yaml:"api_key"`
	MRPPath string `yaml:"mrp_path"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func init() {
	Load(configFile)
}

// Load loads config information into Global
func Load(file string) (AllConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("failed to read config file: %v\n", err)
		return All, err
	}

	err = yaml.Unmarshal(data, &All)
	if err != nil {
		log.Printf("failed to parse config file: %v\n", err)
		return All, err
	}

	Nasa = All.Nasa
	Server = All.Server

	return All, nil
}
