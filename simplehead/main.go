package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("http://localhost")
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)

}
