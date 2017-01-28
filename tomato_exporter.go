package main

import (
	"fmt"
	"net/http"
	"strings"
	"io/ioutil"
	"regexp"
)

type Source struct {
	Download string `json:"rx"`
	Upload string `json:"tx"`
}

type Bandwidth struct {
	Eth0 Source `json:"eth0"`
	Eth1 Source `json:"eth1"`
	Eth2 Source `json:"eth2"`
	Vlan1 Source `json:"vlan1"`
	Vlan2 Source `json:"vlan2"`
	Br0 Source `json:"br0"`
}

func quoteme(b []byte) []byte {
	s := []byte("\"")
	b = append(s, b...)
	b = append(b, s...)
	return b
}

func handlerMain(w http.ResponseWriter, r *http.Request) {

	// ------------ //
	// get the data //
	// ------------ //

	// create the request //
	body := strings.NewReader(`exec=netdev&_http_id=TIDa6f69305333e3371`)
	req, err := http.NewRequest("POST", "http://192.168.10.1/update.cgi", body)
	// set the headers //
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	// do the request //
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error", err)
	}
	defer resp.Body.Close()
	responsebody, _ := ioutil.ReadAll(resp.Body)

	// ---------------------------------- //
	// format the bad json into good json //
	// ---------------------------------- //

	// create the quotinator3000 //
	quotinator3000, _ := regexp.Compile("(0x[\\da-f]+)|(rx)|(tx)")

	// use the quotinator3000
	responsebody = quotinator3000.ReplaceAllFunc(responsebody, quoteme) // add quotes

	// fix excess formatting (additional space and 's)
	for i := 11; i < len(responsebody)-2; i++ {
		if responsebody[i] == 39 { // if it's an '
			responsebody[i] = 34 // make it a "
		}
		responsebody[i-1] = responsebody[i] // move everything from 11 onwards to remove the space
	}

	// slice to get rid of the \n\nnetdata= and };;
	responsebody = responsebody[9:len(responsebody)-3]

	// ------------------ //
	// unmarshal the json //
	// ------------------ //

	fmt.Fprint(w, string(responsebody))

}

func handlerTrash(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/bandwidth", handlerMain)
	http.HandleFunc("/", handlerTrash)
	http.ListenAndServe(":8080", nil)
}