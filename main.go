package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)
	router.Run(":3000")
}

type Metrics struct {
	Report string `json:"report_date"`
	Wind   string `json:"wind"`
	Haw    string `json:"haw"`
	Face   string `json:"face"`
	Scale  string `json:'scale'`
	Note   string `json:"note,omitempty"`
}

func index(c *gin.Context) {
	urls := []string{"north", "south", "east", "west"}

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
