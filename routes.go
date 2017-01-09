package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/sortutil"
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
	Today []Metrics
	Week  []Metrics
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
		reports = append(reports, Report{res.url, data[0:1], data[1:5]})
		sortutil.DescByField(reports, "Coast")
	}

	c.HTML(200, "index.templ.html", gin.H{
		"title":   "Dakine News",
		"reports": reports,
	})
}
