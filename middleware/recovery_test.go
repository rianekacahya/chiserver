package middleware

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestRecovery(t *testing.T) {
	e := chi.NewRouter()
	e.Use(Recovery)
	e.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		panic("error")
	})

	ts := httptest.NewServer(e)
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/ping")
	defer resp.Body.Close()
	if err != nil {
		t.Fatal("Did not expect http.Get to fail")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, `{"message":"Internal Server Error"}`, string(body))
}
