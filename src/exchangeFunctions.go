package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
)

func getExchange(currency string) string {
	var positive string
	var negative string

	response, err := http.Get("http://www.apilayer.net/api/live?access_key=b2e3d360a5c775a403d9ddff35e33cbd&format=1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	value := gjson.GetBytes(responseData, "quotes."+currency+"")

	positive = ":dollar: The exchange rate from " + currency + " = " + value.String()
	negative = "Please check your currency format:exclamation:"

	if gjson.GetBytes(responseData, "quotes."+currency+"").Exists() {
		return (positive)
	} else {
		return (negative)
	}
}
