package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)

	requrest, _ := http.NewRequest("GET", "http://localhsot:18888/upgrade", nil)
	requrest.Header.Set("Connection", "Upgrade")
	requrest.Header.Set("Upgrade", "MyProtocol")
	err = requrest.Write(conn)
	if err != nil {
		panic(err)
	}

	resp, err := http.ReadResponse(reader, requrest)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
	log.Println("Headers: ", resp.Header)

	counter := 10
	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Println("<-", string(bytes.TrimSpace(data)))
		fmt.Fprintf(conn, "%d\n", counter)
		fmt.Println("->", counter)
		counter--

	}
}
