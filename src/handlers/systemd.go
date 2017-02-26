package handlers

import (
	"net/http"
	"../config"
	"os/exec"
	"fmt"
)

func Systemd(w http.ResponseWriter, r *http.Request) {

	conf := config.GetConfig()

	for _,value := range conf.ModSystemd.Services {
		fmt.Println(value)
		output, err := exec.Command("/bin/sh", "-c", "sudo systemctl status "+value).Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprint(w, string(output)+"\n")
	}
}
