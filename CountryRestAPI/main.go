package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"
)

var data map[string]Country

var interval = 10 * time.Minute

var testingMode bool

func init() {
	testingMode = true
	_, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// @title Country API
// @version 1.0
// @description Country Operation(post, get, list, delete)

// @contact.name Gamze Ünal Çoban
// @contact.email u.gamze0@gmail.com

// @host localhost:5003
// @BasePath /
func main() {

	go backgroundTask()

	data = make(map[string]Country)
	readFile, _ := ioutil.ReadFile("countries.json")

	json.Unmarshal(readFile, &data)
	http.HandleFunc("/countries", getCountries)
	http.HandleFunc("/countries/", getCountry)
	http.HandleFunc("/newcountry/", postCountry)
	http.HandleFunc("/countries/delete", delCountries)

	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":5003", logHandler(http.DefaultServeMux)))
}

func backgroundTask() {
	ticker := time.NewTicker(interval)
	for _ = range ticker.C {
		jsonData, _ := json.MarshalIndent(data, "", " ")
		_ = ioutil.WriteFile("countries.json", jsonData, 0644)
	}
}
