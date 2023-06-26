package middleware

import (
	"net/http"

	"github.com/aleale2121/go-eccomerce-app/sessions"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		_, ok := session.Values["USERID"]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
