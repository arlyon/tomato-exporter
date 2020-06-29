package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/arlyon/tomato_exporter/pkg/handlers"

	c "github.com/arlyon/tomato_exporter/configs"
)

func main() {
	var configPath = flag.Arg(0)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Please provide a config: %s ./path-to-config.yaml\n", os.Args[0])
	}

	flag.Parse()

	if configPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	c.LoadConfig(os.Args[1])

	if c.Conf.Modules.ModBandwidth != nil {
		if c.Conf.Modules.ModBandwidth.Slug == "" {
			log.Fatal("Cannot bind mod_bandwidth to route /")
		}
		http.HandleFunc("/"+c.Conf.Modules.ModBandwidth.Slug, handlers.Bandwidth)
	}
	if c.Conf.Modules.ModSystemd != nil {
		if c.Conf.Modules.ModSystemd.Slug == "" {
			log.Fatal("Cannot bind mod_systemd to route /")
		}
		http.HandleFunc("/"+c.Conf.Modules.ModSystemd.Slug, handlers.Systemd)
	}

	http.HandleFunc("/", handlerBase)
	fmt.Println(fmt.Sprintf("Now listening on http://%s:%d", c.Conf.IP, c.Conf.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", c.Conf.Port), nil)
	if err != nil {
		fmt.Println("Could not bind to port! Shutting down: " + err.Error())
	}
}

func handlerBase(w http.ResponseWriter, r *http.Request) {
	v := "<h1>Tomato Exporter</h1>\n"

	if c.Conf.Modules.ModBandwidth != nil {
		name := c.Conf.Modules.ModBandwidth.Slug
		v += fmt.Sprintf("<ul><a href=\"/%s\">%s</a></ul>", strings.ToLower(fmt.Sprint(name)), name)
	}

	if c.Conf.Modules.ModSystemd != nil {
		name := c.Conf.Modules.ModSystemd.Slug
		v += fmt.Sprintf("<ul><a href=\"/%s\">%s</a></ul>", strings.ToLower(fmt.Sprint(name)), name)
	}

	fmt.Fprint(w, v)
}
