package handlers

import (
	"net/http"
	"../config"
	"os/exec"
	"fmt"
	"strings"
	"github.com/fatih/structs"
	"encoding/json"
)

type Status struct {
	Loaded bool
	Active bool
	Substate string
}

func Systemd(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	systemd, err := exec.Command("/bin/sh", "-c", "sudo systemctl list-units | grep '"+strings.Join(conf.ModSystemd.Services, "\\|")+"'").Output()
	if err != nil {
		fmt.Println(err)
	}

	response := make(map[string]interface{})

	for _,element := range strings.Split(string(systemd[1:len(systemd)-1]), "\n") {
		args := strings.Fields(element)

		isloaded := false
		if args[1] == "loaded" {
			isloaded = true
		}

		isactive := false
		if args[2] == "active" {
			isactive = true
		}

		service := structs.Map(Status{isloaded, isactive, args[3]})
		response[args[0]] = service
	}

	responsebody, _ := json.Marshal(response)
	responsestring := string(responsebody)

	fmt.Fprint(w, responsestring)
}
