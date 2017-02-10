package handlers

import (
	"strings"
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
	"../config"
	"encoding/json"
	"encoding/base64"
)

func Bandwidth(w http.ResponseWriter, r *http.Request) {

	conf := config.GetConfig()

	// ------------ //
	// get the data //
	// ------------ //

	// create the request //
	body := strings.NewReader(`exec=netdev&_http_id=`+conf.HttpId)
	req, err := http.NewRequest("POST", "http://"+conf.Ip+"/update.cgi", body)
	// authenticate by converting the username and password to base 64 //
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",conf.Username,conf.Password)))
	req.Header.Set("Authorization", "Basic "+auth)

	// do the request //
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error", err)
	}
	defer resp.Body.Close()
	responsebody, _ := ioutil.ReadAll(resp.Body)

	// fix excess formatting (additional space and 's)
	for i := 11; i < len(responsebody)-2; i++ {
		if responsebody[i] == 39 { // if it's an '
			responsebody[i] = 34 // make it a "
		}
		responsebody[i-1] = responsebody[i] // move everything from 11 onwards to remove the space
	}

	// slice to get rid of the \n\nnetdata= and };;
	responsebody = responsebody[9:len(responsebody)-3]

	// ---------------------------------- //
	// format the bad json into good json //
	// ---------------------------------- //

	// create and use the quotinator3000 //
	quotinator3000, _ := regexp.Compile("(rx)|(tx)")
	responsestring := quotinator3000.ReplaceAllStringFunc(string(responsebody), quoteme) // add quotes

	// create and use the dehexinator2000 //
	dehexinator, _ := regexp.Compile("(0x[\\da-f]+)")
	responsestring = dehexinator.ReplaceAllStringFunc(responsestring, dehex) // add quotes

	if conf.ModBandwidth.Interfaces[0] != "all" {
		var data map[string]interface{}
		response := make(map[string]interface{})
		err = json.Unmarshal([]byte(responsestring), &data)
		for _,value := range conf.ModBandwidth.Interfaces {
			for key,data := range data {
				if value == key {
					response[value] = data
				}
			}
		}
		responsebody, _ = json.Marshal(response)
		responsestring = string(responsebody)
	}

	// ------------- //
	// send it away! //
	// ------------- //

	fmt.Fprint(w, responsestring)

}