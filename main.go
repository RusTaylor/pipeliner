package main

import (
	"pipeliner/config"
	"pipeliner/migration"
	"pipeliner/server"
)

func main() {
	config.SetConfig()
	migration.Migrate()
	//todo проверка на наличие баз данных и если нет, то создать
	server.StartHTTPServer(config.Config.Server.Host, config.Config.Server.Port)
}
