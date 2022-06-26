package wlog

import "net/http"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		New().Infof("Request received for %s from IP: %s", r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
