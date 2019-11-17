package middleware

import (
	"errors"
	"net/http"
	"github.com/rianekacahya/chiserver/response"
)

var(
	errorString = errors.New("Internal Server Error")
)

func Recovery(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response.Error(w, errorString)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
