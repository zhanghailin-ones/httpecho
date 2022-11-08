// httpecho starts a http server which prints all incoming requests to the log and responses with
// code 200
// Install: go install github.com/zhanghailin-ones/httpecho@latest
// Usage: httpecho -port 8880
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Printf("httputil.DumpRequest(): %v", err)
		}
		log.Println(string(b))
	})

	log.Printf("listeing on port: %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
