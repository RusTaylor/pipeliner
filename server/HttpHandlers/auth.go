package HttpHandlers

import (
	"html/template"
	"log"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/login.html")

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
}

func Auth(w http.ResponseWriter, r *http.Request) {
	requestMethod := r.Method
	if requestMethod != "POST" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Page not found!"))
		return
	}

}
