package server

import (
	"log"
	"net/http"
)

func StartHTTPServer(domain string, port string) {
	setRoutes()
	log.Println("HTTP Server Start")
	err := http.ListenAndServe(domain+":"+port, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func StartHTTPSServer(domain string, port string, certFile string, keyFile string) {
	setRoutes()
	log.Println("HTTPS Server Start")
	err := http.ListenAndServeTLS(domain+":"+port, certFile, keyFile, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
