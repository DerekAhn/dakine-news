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

type Report struct {
	Coast string
	Data  []Metrics
}

func index(c *gin.Context) {
	urls := []string{"north", "west", "south", "east"}

	responses := asyncHttpGets(urls)
	reports := make([]Report, 0)

	for _, res := range responses {
		var data []Metrics
		if err := json.Unmarshal(res.body, &data); err != nil {
			log.Fatalln("Error decoing JSON", err)
		}
		reports = append(reports, Report{res.url, data})
	}

	c.HTML(200, "index.templ.html", gin.H{
		"title":   "Dakine News",
		"reports": reports,
	})
}
