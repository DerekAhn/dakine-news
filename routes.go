package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type Metrics struct {
	Report string `json:"report_date"`
	Wind   string `json:"wind"`
	Haw    string `json:"haw"`
	Face   string `json:"face"`
	Scale  string `json:'scale'`
	Note   string `json:"note,omitempty"`
}

func index(c *gin.Context) {
	urls := []string{"north", "west", "south", "east"}

	responses := asyncHttpGets(urls)
	reports := make(map[string][]Metrics)

	for _, res := range responses {
		var report []Metrics
		if err := json.Unmarshal(res.body, &report); err != nil {
			log.Fatalln("Error decoing JSON", err)
		}

		reports[res.url] = report
	}

	c.HTML(200, "index.templ.html", gin.H{
		"world":   "Gopher",
		"reports": reports,
	})
}
