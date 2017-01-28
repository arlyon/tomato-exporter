package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"os"
	"reflect"
	"unsafe"

	"./config"
	"./handlers"
)

var c = config.Config{}

func quoteme(b []byte) []byte {
	s := []byte("\"")
	b = append(s, b...)
	b = append(b, s...)
	return b
}

func handlerFavicon(w http.ResponseWriter, r *http.Request) {}

func handlerBase(w http.ResponseWriter, r *http.Request) {
	attributes := int(unsafe.Sizeof(c.Modules))

	v := "<h1>Tomato Exporter</h1>\n"

	for i:=0;i<attributes;i++ {
		name := reflect.ValueOf(c.Modules).Type().Field(i).Name
		if reflect.ValueOf(c.Modules).Field(0).Bool() == true {
			v += fmt.Sprintf("<ul><a href=\"/%s\">%s</a></ul>",strings.ToLower(name),name)
		}
	}

	fmt.Fprint(w, v)
}

func main() {

	// open the config file //
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// parse the config file //
	err = json.Unmarshal(configFile, &c)
	if err != nil {
		fmt.Println("Bad formatting in config: ", err)
		os.Exit(2)
	}

	// start the web server //
	if c.Modules.Bandwidth == true {
		http.HandleFunc("/bandwidth", handlers.Bandwidth)
	}
	http.HandleFunc("/favicon.ico", handlerFavicon)
	http.HandleFunc("/", handlerBase)
	fmt.Printf("Now listening on port %d.", c.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil)
}