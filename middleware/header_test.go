package middleware

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeader(t *testing.T) {
	e := chi.NewRouter()
	e.Use(Header)
	e.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	ts := httptest.NewServer(e)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal("Did not expect http.Get to fail")
	}
	assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))
	assert.Equal(t, "off", resp.Header.Get("X-DNS-Prefetch-Control"))
	assert.Equal(t, "DENY", resp.Header.Get("X-Frame-Options"))
	assert.Equal(t, "max-age=5184000; includeSubDomains", resp.Header.Get("Strict-Transport-Security"))
	assert.Equal(t, "noopen", resp.Header.Get("X-Download-Options"))
	assert.Equal(t, "1; mode=block", resp.Header.Get("X-XSS-Protection"))
}
