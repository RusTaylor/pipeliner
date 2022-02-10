package main

import (
	"pipeliner/server"
)

func main() {
	//todo проверка на наличие баз данных и если нет, то создать
	server.StartHTTPServer("pipeliner", "8080")
}
