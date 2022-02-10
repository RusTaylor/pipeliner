package HttpHandlers

import (
	"html/template"
	"log"
	"net/http"
	"pipeliner/server/session"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if !session.IsLogin(r) {
		w.Header().Set("Cache-Control", "no-cache, private") // HTTP 1.1.
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		log.Println(err)
	}
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Println(err)
	}

	user := session.GetUser(cookie.Value)

	err = tmpl.Execute(w, user)

	if err != nil {
		log.Println(err)
	}
}
