package HttpHandlers

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"pipeliner/database"
	"pipeliner/server/session"
	"pipeliner/server/user"
)

type UserFromDb struct {
	Id       int
	Login    string
	Name     string
	Password string
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if session.IsLogin(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("public/login.html")

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	w.Header().Set("Cache-Control", "no-cache, private") // HTTP 1.1.
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

	db, err := database.GetDb()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	defer db.Close()

	r.ParseForm()
	login := r.FormValue("login")
	password := r.FormValue("password")

	if login == "" || password == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var userFromDb UserFromDb
	err = db.QueryRow("SELECT id,login,name,password FROM \"user\"").Scan(
		&userFromDb.Id,
		&userFromDb.Login,
		&userFromDb.Name,
		&userFromDb.Password)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(password)); err == nil {
		session.SetSession(w, user.User{Id: userFromDb.Id, Login: userFromDb.Login, Name: userFromDb.Name})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return
}
