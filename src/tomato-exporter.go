package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"./handlers"

	c "./config"
)

func handlerFavicon(w http.ResponseWriter, r *http.Request) {}

func handlerBase(w http.ResponseWriter, r *http.Request) {
	v := "<h1>Tomato Exporter</h1>\n"

	if c.Conf.ModBandwidth != nil {
		name := c.Conf.ModBandwidth.Slug
		v += fmt.Sprintf("<ul><a href=\"/%s\">%s</a></ul>", strings.ToLower(fmt.Sprint(name)), name)
	}

	if c.Conf.ModSystemd != nil {
		name := c.Conf.ModSystemd.Slug
		v += fmt.Sprintf("<ul><a href=\"/%s\">%s</a></ul>", strings.ToLower(fmt.Sprint(name)), name)
	}

	fmt.Fprint(w, v)
}

func main() {
	if len(os.Args) != 2 {
		println("usage: " + os.Args[0] + " ./path-to-config.conf")
		os.Exit(1)
	}

	c.LoadConfig(os.Args[1])

	if c.Conf.ModBandwidth != nil {
		if c.Conf.ModBandwidth.Slug == "" {
			log.Fatal("Cannot bind mod_bandwidth to route /")
		}
		http.HandleFunc("/"+c.Conf.ModBandwidth.Slug, handlers.Bandwidth)
	}
	if c.Conf.ModSystemd != nil {
		if c.Conf.ModSystemd.Slug == "" {
			log.Fatal("Cannot bind mod_systemd to route /")
		}
		http.HandleFunc("/"+c.Conf.ModSystemd.Slug, handlers.Systemd)
	}

	// start the web server //
	http.HandleFunc("/favicon.ico", handlerFavicon)
	http.HandleFunc("/", handlerBase)
	fmt.Println(fmt.Sprintf("Now listening on http://%s:%d", c.Conf.IP, c.Conf.Port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", c.Conf.Port), nil)
	if err != nil {
		fmt.Println("Could not bind to port! Shutting down.")
	}
}
