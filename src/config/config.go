package config

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

var config = Config{}

func SetConfig() Config {
	// open the config file //
	configFile, err := ioutil.ReadFile("tomato-exporter.conf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// parse the config file //
	err = json.Unmarshal(configFile, &config)
	fmt.Println(config)
	if err != nil {
		fmt.Println("Bad formatting in config: ", err)
		os.Exit(2)
	}
	return config
}

func GetConfig() Config {
	return config;
}

type Config struct {
	Port int       			`json:"hosting_port"`
	HttpId string       		`json:"http_id"`
	Username string       		`json:"router_username"`
	Password string       		`json:"router_password"`
	Ip string 			`json:"router_ip"`
	EnabledModules Modules 		`json:"enabled_modules"`
	ModBandwidth ModuleSettings    	`json:"mod_bandwidth"`
}

type Modules struct {
	ModBandwidth bool      	`json:"mod_bandwidth"`
}

type ModuleSettings struct {
	Interfaces []string 		`json:"interfaces"`
	Slug string 			`json:"slug"`
}