package middleware

import (
	"net/http"
)

func Headers(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Protects from MimeType Sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")
		// Prevents browser from prefetching DNS
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		// Denies website content to be served in an iframe
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Strict-Transport-Security", "max-age=5184000; includeSubDomains")
		// Prevents Internet Explorer from executing downloads in site's context
		w.Header().Set("X-Download-Options", "noopen")
		// Minimal XSS protection
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}