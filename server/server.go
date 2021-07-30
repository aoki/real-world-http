package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/k0kubun/pp"
)

func handlerDigest(w http.ResponseWriter, r *http.Request) {
	pp.Println("URL: %s\n", r.URL.String())
	pp.Println("Query: %v\n", r.URL.Query())
	pp.Println("Proto: %s\n", r.Proto)
	pp.Println("Method: %s\n", r.Method)
	pp.Println("Header: %v\n", r.Header)

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("--body--\n%s\n", string(body))
	if _, ok := r.Header["Authorization"]; !ok {
		w.Header().Add("WWW-Authenticate", `Digest realm="Secret Zone", nonce="TgLc25U2BQA=f510fkjwi3aw3jfaw3r3ra", algorithm=MD5, qop="auth"`)
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Fprintf(w, "<html><body>secret page</body></html>\n")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Set-Cookie", "VISIT=TRUE")
	pp.Println("Query: %v\n", r.URL.Query())

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

func handleSlow(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		log.Print(i)
	}

	fmt.Fprintf(w, "<html><body>10</body></html>\n")
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/digest", handlerDigest)
	http.HandleFunc("/slow", handleSlow)
	log.Println("start http listening :80")
	httpServer.Addr = ":80"
	log.Println(httpServer.ListenAndServe())
}
