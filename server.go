package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")

	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body>2回目以降</body></html>\n")
	} else {
		fmt.Fprintf(w, "<html><body>hello 初訪問</body></html>\n")
	}
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	log.Println("start http listening :80")
	httpServer.Addr = ":80"
	log.Println(httpServer.ListenAndServe())
}
