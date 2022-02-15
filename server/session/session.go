package session

import (
	"math/rand"
	"net/http"
	"pipeliner/server/user"
	"strings"
	"sync"
	"time"
)

var sessions = Sessions{List: make(map[string]Session)}

type Session struct {
	User    user.User
	DateDie time.Time
}

type Sessions struct {
	List map[string]Session
	sync.RWMutex
}

func IsLogin(request *http.Request) bool {
	cookie, _ := request.Cookie("token")
	if cookie == nil {
		return false
	}

	sessions.Lock()
	defer sessions.Unlock()
	if val, isset := sessions.List[cookie.Value]; isset {
		currentDate := time.Now()

		if currentDate.Before(val.DateDie) {
			return true
		}

		delete(sessions.List, cookie.Value)
		return false
	}

	return false
}

func SetSession(w http.ResponseWriter, user user.User) {
	token := getToken()
	dateDie := time.Now().Add(9 * time.Hour)

	sessions.Lock()
	sessions.List[token] = Session{User: user, DateDie: dateDie}
	sessions.Unlock()

	cookie := &http.Cookie{
		Name:  "token",
		Value: token,
	}

	http.SetCookie(w, cookie)
}

func GetUser(sessionToken string) *user.User {
	sessions.RLock()
	defer sessions.RUnlock()

	if val, isset := sessions.List[sessionToken]; isset {
		return &val.User
	}

	return nil
}

func getToken() string {
	token := generateToken()

	sessions.RLock()
	defer sessions.RUnlock()
	if _, isset := sessions.List[token]; isset {
		return getToken()
	}

	return token
}

func generateToken() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
