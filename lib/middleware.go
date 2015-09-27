package passwordless

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/justinas/nosurf"
	"net/http"
)

// CsrfMiddleware adds CSRF support via nosurf.
func CsrfMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var token string
	var passed bool

	// nosurf disposes of the token as soon as it calls the http.Handler you provide...
	// in order to use it as negroni middleware, pull out token and dispose of it ourselves
	csrfHandler := nosurf.NewPure(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		token = nosurf.Token(r)
		passed = true
	}))
	csrfHandler.ServeHTTP(w, r)

	// csrf passed
	if passed {
		context.Set(r, "csrf_token", token)
		next(w, r)
		context.Delete(r, "csrf_token")
	}
}

// UserMiddleware checks for the User in the session and adds them to the request context if they exist.
func UserMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Debug("in UserMiddleware...")
	s := GetSession(r)
	log.Debug("is new ... ", s.IsNew)
	if id, ok := s.Values[sessionUser]; ok {
		log.Debug("id is", id)
		if user, err := dbmap.Get(User{}, id.(int64)); err == nil {
			SetContextUser(user.(*User), r)
		} else {
			log.Debug("dbmap had an error...")
		}
	} else {
		log.Debug("UserMiddleware ... s.Values[sessionUser] not ok...")
		log.Debug("UserMiddleware - s values...", s.Values)
	}
	next(w, r)
}

// LoginRequiredMiddleware ensures a User is logged in, otherwise redirects them to the login page.
func LoginRequiredMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if !IsLoggedIn(r) {
		log.Debug("user not logged in...redirecting...")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	next(w, r)
}
