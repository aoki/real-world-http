package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm("http://localhost", values)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
}
