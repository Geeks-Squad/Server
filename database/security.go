package database

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// BasicAuth Middleware with next handler in chain
type BasicAuth struct {
	Next http.Handler
}

func (b BasicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	b.Next.ServeHTTP(w, r)
	return

	authorizationArray := r.Header["Authorization"]

	if len(authorizationArray) > 0 {
		authorization := strings.TrimSpace(authorizationArray[0])
		credentials := strings.Split(authorization, " ")

		if len(credentials) != 2 || credentials[0] != "Basic" {
			unauthorized(w)
			return
		}

		authstr, err := base64.StdEncoding.DecodeString(credentials[1])
		if err != nil {
			unauthorized(w)
			return
		}

		userpass := strings.Split(string(authstr), ":")
		if len(userpass) != 2 {
			unauthorized(w)
			return
		}

		if userpass[0] == "foo" && userpass[1] == "bar" {

		} else {
			unauthorized(w)
		}
	} else {
		unauthorized(w)
	}
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Authorization", "No")
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
