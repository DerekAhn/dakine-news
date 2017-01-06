package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Metrics struct {
	Report string `json:"report_date"`
	Wind   string `json:"wind"`
	Haw    string `json:"haw"`
	Face   string `json:"face"`
	Scale  string `json:'scale'`
	Note   string `json:"note"`
}

var urls = []string{
	"north",
	"south",
	"east",
	"west",
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.Run(":3000")
}

func index(c *gin.Context) {
	responses := asyncHttpGets(urls)
	reports := make(map[string][]Metrics)

	for _, res := range responses {
		var report []Metrics
		if err := json.Unmarshal(res.body, &report); err != nil {
			log.Fatalln("Error decoing JSON", err)
		}
		reports[res.url] = report
	}

	c.JSON(200, reports)
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

type HttpResponse struct {
	url  string
	body []byte
	err  error
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls))
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			data, err := fetch(url)

			ch <- &HttpResponse{url, data, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		default:
			time.Sleep(5e1)
		}
	}

	return responses
}
