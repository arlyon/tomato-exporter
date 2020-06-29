package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Conf contains the configs for this app
var Conf = Config{}

// LoadConfig does this
func LoadConfig(path string) Config {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(configFile, &Conf)
	if err != nil {
		fmt.Println("Bad formatting in config: ", err)
		os.Exit(2)
	}
	return Conf
}

// Config describes the options available to the program
type Config struct {
	Port    int     `yaml:"port"`
	IP      string  `yaml:"ip"`
	Modules modules `yaml:"modules"`
}

type modules struct {
	ModBandwidth *modBandwidth `yaml:"mod_bandwidth"`
	ModSystemd   *modSystemD   `yaml:"mod_systemd"`
}

type modSystemD struct {
	Slug     string   `yaml:"slug"`
	Services []string `yaml:"services"`
}

type modBandwidth struct {
	Slug       string   `yaml:"slug"`
	Interfaces []string `yaml:"interfaces"`
	HTTPID     string   `yaml:"http_id"`
	Username   string   `yaml:"admin_username"`
	Password   string   `yaml:"admin_password"`
	IP         string   `yaml:"router_ip"`
}
