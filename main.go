package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type Metrics struct {
	Report string `json:"report_date"`
	Wind   string `json:"wind"`
	Haw    string `json:"haw"`
	Face   string `json:"face"`
	Scale  string `json:'scale'`
	Note   string `json:"note"`
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.Run(":3000")
}

func index(c *gin.Context) {

	data, err := fetch("east")
	checkErr(err, "API fetch failed")

	var east []Metrics
	if err := json.Unmarshal(data, &east); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	c.JSON(200, east)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func fetch(url string) ([]byte, error) {
	const API string = "http://www.surfnewsnetwork.com/json-api/"

	resp, err := http.Get(API + url)
	checkErr(err, "Error fetching API")

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
