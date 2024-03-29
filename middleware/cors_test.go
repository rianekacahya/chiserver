package middleware

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	e := chi.NewRouter()
	e.Use(CORS)
	e.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	ts := httptest.NewServer(e)
	defer ts.Close()
	var cl http.Client
	req, _ := http.NewRequest("OPTIONS", ts.URL+"/ping", nil)
	resp, _ := cl.Do(req)
	assert.Equal(t, "86400", resp.Header.Get("Access-Control-Max-Age"))
	assert.Equal(t, "POST,GET,PUT,DELETE,PATCH,HEAD", resp.Header.Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "true", resp.Header.Get("Access-Control-Allow-Credentials"))
}
