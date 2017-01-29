package handlers

import (
	"fmt"
	"strconv"
)

func quoteme(input string) string {
	quote := "\""
	output := fmt.Sprintf(quote+input+quote)
	return output
}

func dehex(b string) string {
	number, err := strconv.ParseInt(string(b[2:]), 16, 0)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%v", number);
}