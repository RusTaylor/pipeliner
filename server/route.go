package server

import (
	"net/http"
	"pipeliner/server/HttpHandlers"
)

func setRoutes() {
	http.HandleFunc("/login", HttpHandlers.LoginPage)
	http.HandleFunc("/auth", HttpHandlers.Auth)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.HandleFunc("/favicon.ico", faviconHandler)
	//Роуты требующие проверки на логин
	http.HandleFunc("/", HttpHandlers.Index)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/favicon.ico")
}
