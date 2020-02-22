package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	c "github.com/arlyon/tomato_exporter/configs"
	"github.com/fatih/structs"
)

// Status shows the status of a service.
type Status struct {
	Loaded   int
	Active   int
	Substate string
}

// Systemd is a handler designed to pump systemd stats for prometheus
func Systemd(w http.ResponseWriter, r *http.Request) {
	systemd, err := exec.Command("/bin/sh", "-c", "sudo systemctl list-units | grep '"+strings.Join(c.Conf.ModSystemd.Services, "\\|")+"'").Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if len(systemd) == 0 {
		fmt.Fprintf(w, "{}")
		return
	}

	response := make(map[string]interface{})
	for _, element := range strings.Split(string(systemd[1:len(systemd)-1]), "\n") {
		args := strings.Fields(element)

		isloaded := 0
		if args[1] == "loaded" {
			isloaded = 1
		}

		isactive := 0
		if args[2] == "active" {
			isactive = 1
		}

		service := structs.Map(Status{isloaded, isactive, args[3]})
		response[args[0]] = service
	}

	responsebody, _ := json.Marshal(response)
	responsestring := string(responsebody)

	fmt.Fprint(w, responsestring)
}
