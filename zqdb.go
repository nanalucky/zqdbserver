package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/csv"
	"strings"
	"log"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	configData,_ := ioutil.ReadFile("config.txt")
	fmt.Println(string(configData))

	accountData,_ := ioutil.ReadFile("account.txt")
	//fmt.Print(string(accountData))
	r := csv.NewReader(strings.NewReader(string(accountData)))
	accountRecords, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Print(accountRecords)
	for i:=0; i < len(accountRecords); i++	{
		for j:=0; j < len(accountRecords[i]); j++{
			fmt.Println(accountRecords[i][j])
		}
	}


	http.HandleFunc("/config/", configHandler)
	http.ListenAndServe(":8080", nil)
}