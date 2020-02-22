package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	c "../config"
)

// Bandwidth is a handler to publish bandwidth metrics in prometheus format
func Bandwidth(w http.ResponseWriter, r *http.Request) {
	// create the request //
	body := strings.NewReader(`exec=netdev&_http_id=` + c.Conf.ModBandwidth.HTTPID)
	req, err := http.NewRequest("POST", "http://"+c.Conf.ModBandwidth.IP+"/update.cgi", body)
	if err != nil {
		http.Error(w, "error creating request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// authenticate by converting the username and password to base 64 //
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.Conf.ModBandwidth.Username, c.Conf.ModBandwidth.Password)))
	req.Header.Set("Authorization", "Basic "+auth)

	// do the request //
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "error sending request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	responsebody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "error reading request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// fix excess formatting (additional space and 's)
	for i := 11; i < len(responsebody)-2; i++ {
		if responsebody[i] == 39 { // if it's an '
			responsebody[i] = 34 // make it a "
		}
		responsebody[i-1] = responsebody[i] // move everything from 11 onwards to remove the space
	}

	// slice to get rid of the \n\nnetdata= and };;
	responsebody = responsebody[9 : len(responsebody)-3]

	// ---------------------------------- //
	// format the bad json into good json //
	// ---------------------------------- //

	// create and use the quotinator3000 //
	quotinator3000, _ := regexp.Compile("(rx)|(tx)")
	responsestring := quotinator3000.ReplaceAllStringFunc(string(responsebody), quoteme)

	// create and use the dehexinator2000 //
	dehexinator, _ := regexp.Compile("(0x[\\da-f]+)")
	responsestring = dehexinator.ReplaceAllStringFunc(responsestring, dehex)

	if len(c.Conf.ModBandwidth.Interfaces) != 0 {
		var data map[string]interface{}
		response := make(map[string]interface{})
		err = json.Unmarshal([]byte(responsestring), &data)
		for _, value := range c.Conf.ModBandwidth.Interfaces {
			for key, data := range data {
				if value == key {
					response[value] = data
				}
			}
		}
		responsebody, _ = json.Marshal(response)
		responsestring = string(responsebody)
	}

	fmt.Fprint(w, responsestring)
}
