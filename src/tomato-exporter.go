package main

import (
	"fmt"
	"net/http"
	"./config"
	"./handlers"
	"./github.com/fatih/structs"
	"strings"
)

var conf = config.Config{}

func handlerFavicon(w http.ResponseWriter, r *http.Request) {}

func handlerBase(w http.ResponseWriter, r *http.Request) {
	confmap := structs.Map(conf)

	v := "<h1>Tomato Exporter</h1>\n"
	enabled_modules := confmap["EnabledModules"].(map[string]interface{})

	for key, value := range enabled_modules {
		if value == true {
			name := confmap[key].(map[string]interface{})["Slug"]
			v += fmt.Sprintf("<ul><a href=\"/%s\">%s</a></ul>",strings.ToLower(fmt.Sprint(name)),name)
		}
	}

	fmt.Fprint(w, v)
}

func main() {

	conf = config.SetConfig()

	// start the web server //
	if conf.EnabledModules.ModBandwidth == true {
		http.HandleFunc("/"+conf.ModBandwidth.Slug, handlers.Bandwidth)
	}
	http.HandleFunc("/favicon.ico", handlerFavicon)
	http.HandleFunc("/", handlerBase)
	fmt.Println("Now listening on port", conf.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), nil)
	if err != nil {
		fmt.Println("Port is in use! Shutting down.")
	}
}