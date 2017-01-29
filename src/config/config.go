package config

type Config struct {
	Port int 	`json:"port"`
	Modules Module `json:"enabled_modules"`
}

type Module struct {
	Bandwidth bool 	`json:"bandwidth"`
}