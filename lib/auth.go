package passwordless

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

const (
	contextUser string = "User"
	sessionUser string = "_user_id"
)

var store *sessions.CookieStore
var secretKey []byte

// GetSession returns the session for the site.
func GetSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "site")
	if err != nil {
		log.Error("problem decoding session ... ", err)
	}
	return session
}

// GetContextUser returns the User for the given request context or nil.
func GetContextUser(r *http.Request) *User {
	log.Debug("context contents are...", context.GetAll(r))
	if user, ok := context.GetOk(r, contextUser); ok {
		return user.(*User)
	}
	log.Debug("in GetContextUser...no user!")
	return nil
}

// SetContextUser stores the given user in the request context.
func SetContextUser(user *User, r *http.Request) {
	log.Debug("setting context!")
	context.Set(r, contextUser, user)
}

// Login adds the User's id to the session.
func Login(u *User, w http.ResponseWriter, r *http.Request) {
	log.Debug("in login...", u.Id)
	s := GetSession(r)
	s.Values[sessionUser] = u.Id
	err := s.Save(r, w)
	if err != nil {
		log.Debug("...something went wrong saving!!!...", err)
	}
	log.Debug("in login after session save...", s.Values)
	//SetContextUser(u, r)
}

// Logout removes the User from their session.
func Logout(w http.ResponseWriter, r *http.Request) {
	s := GetSession(r)
	delete(s.Values, sessionUser)
	s.Save(r, w)
}

// IsLoggedIn is a convenience function for checking if a User exists in the request context.
func IsLoggedIn(r *http.Request) bool {
	return GetContextUser(r) != nil
}

func init() {
	//log.Debug("****calling init in auth...")
	authKey := os.Getenv("AUTH_KEY")
	if authKey == "" {
		log.Panic("Missing required environment variable 'AUTH_KEY'.")
	}

	encryptKey := os.Getenv("ENCRYPT_KEY")
	if encryptKey == "" {
		log.Panic("Missing required environment variable 'ENCRYPT_KEY' (16, 24, 32 bytes in length).")
	}

	store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32)) //[]byte(authKey), []byte(encryptKey))

	//	store.Options = &sessions.Options{
	//		Domain:   "localhost",
	//		Path:     "/",
	//		MaxAge:   3600 * 8, // 8 hours
	//		HttpOnly: true,
	//	}

}
