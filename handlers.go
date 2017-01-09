package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

func fetch(url string) ([]byte, error) {
	const API string = "http://www.surfnewsnetwork.com/json-api/"

	resp, err := http.Get(API + url)

	if err != nil {
		log.Fatalln("Error http.GET:", err)
		log.Fatalln("Error failed GET:", API+url, "\nError Status not OK:", resp.StatusCode)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
