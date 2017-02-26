package handlers

import (
	"net/http"
	"../config"
	"os/exec"
	"fmt"
	"strings"
)

func Systemd(w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	output, err := exec.Command("/bin/sh", "-c", "sudo systemctl list-units | grep '"+strings.Join(conf.ModSystemd.Services, "\\|")+"'").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprint(w, string(output)+"\n")
}
