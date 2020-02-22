package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	err = json.Unmarshal(configFile, &Conf)
	if err != nil {
		fmt.Println("Bad formatting in config: ", err)
		os.Exit(2)
	}
	return Conf
}

// Config describes the options available to the program
type Config struct {
	Port         int           `json:"hosting_port"`
	IP           string        `json:"binding_ip"`
	ModBandwidth *modBandwidth `json:"mod_bandwidth"`
	ModSystemd   *modSystemD   `json:"mod_systemd"`
}

type modules struct {
	ModBandwidth bool `json:"mod_bandwidth"`
	ModSystemd   bool `json:"mod_systemd"`
}

type modSystemD struct {
	Slug     string   `json:"slug"`
	Services []string `json:"services"`
}

type modBandwidth struct {
	Slug       string   `json:"slug"`
	Interfaces []string `json:"interfaces"`
	HTTPID     string   `json:"http_id"`
	Username   string   `json:"router_username"`
	Password   string   `json:"router_password"`
	IP         string   `json:"router_ip"`
}
