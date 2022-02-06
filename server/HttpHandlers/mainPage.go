package HttpHandlers

import (
	"html/template"
	"log"
	"net/http"
	"pipeliner/server/session"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if !session.IsLogin(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Println(err)
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err)
	}
}
