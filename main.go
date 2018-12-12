package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

func main() {
	listen := flag.Int("listen", 8080, "The port to listen locally")
	target := flag.String("target", "https://http2.golang.org/", "The URL to proxy")
	key := flag.String("header-key", "", "The key of the injected header")
	value := flag.String("header-value", "", "The value of the injected header")
	flag.Parse()

	url, err := url.Parse(*target)
	if err != nil || url.Host == "" {
		log.Println("Invalid target URL")
		flag.PrintDefaults()
		os.Exit(1)
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(url)

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// Inject custom header
		if *key != "" {
			req.Header.Set(*key, *value)
		}
		reverseProxy.ServeHTTP(rw, req)
	})
	log.Panic(http.ListenAndServe(":"+strconv.Itoa(*listen), nil))
}
