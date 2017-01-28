package config

type Config struct {
	Port int 	`json:"port,string"`
	Modules Module `json:"enabled_modules,string"`
}

type Module struct {
	Bandwidth bool 	`json:"bandwidth"`
}

type Source struct {
	Download string `json:"rx"`
	Upload string `json:"tx"`
}

type Bandwidth struct {
	Eth0 Source `json:"eth0"`
	Eth1 Source `json:"eth1"`
	Eth2 Source `json:"eth2"`
	Vlan1 Source `json:"vlan1"`
	Vlan2 Source `json:"vlan2"`
	Br0 Source `json:"br0"`
}