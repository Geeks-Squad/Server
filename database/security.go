package database

import (
	"net/http"
	"strings"
)

// BasicAuth Middleware with next handler in chain
type BasicAuth struct {
	Next http.Handler
}

func (b BasicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	authorizationArray := r.Header["Authorization"]

	authorization := strings.TrimSpace(authorizationArray[0])

	userpass := strings.Split(authorization, ":")
	if len(userpass) != 2 {
		unauthorized(w)
		return
	}
	if CheckLogin(userpass[0], userpass[1]) {
		b.Next.ServeHTTP(w, r)
	} else {
		unauthorized(w)
	}

}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Authorization", "No")
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
