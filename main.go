package main

import (
	"pipeliner/server"
)

func main() {
	server.StartHTTPServer("pipeliner", "8080")
}
