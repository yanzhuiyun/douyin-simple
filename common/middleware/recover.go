package middleware

import (
	"common/result"
	"net/http"
)

func RecoverFromPanic(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				result.ResponseBadServer(w)
			}
		}()
		next(w, r)
	}
}
