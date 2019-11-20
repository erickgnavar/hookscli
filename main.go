package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func startServer(server string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// basic data
		fmt.Fprintf(os.Stdout, "%s %s %s\n", r.Method, r.URL, r.Proto)

		// body handling
		var stringBuilder strings.Builder
		b := make([]byte, 8)
		for {
			n, err := r.Body.Read(b)
			stringBuilder.WriteString(fmt.Sprintf("%q", b[:n]))
			if err == io.EOF {
				break
			}
		}

		fmt.Fprintf(os.Stdout, "Body: %s\n", stringBuilder.String())

		// headers
		fmt.Fprintf(os.Stdout, "RemoteAddr: %v\n", r.RemoteAddr)
		for key, value := range r.Header {
			fmt.Fprintf(os.Stdout, "%v: %v\n", key, value)
		}
		fmt.Println("")
		io.WriteString(w, "OK")
	})

	http.ListenAndServe(server, nil)
}

func main() {
	logger := log.New(os.Stdout, "HOOKSCLI: ", log.LstdFlags)

	var port int
	var host string

	flag.StringVar(&host, "host", "localhost", "Host")
	flag.IntVar(&port, "port", 8080, "Port")
	flag.Parse()

	server := fmt.Sprintf("%v:%v", host, port)
	logger.Println("Starting server on %v...", server)

	startServer(server)
}
