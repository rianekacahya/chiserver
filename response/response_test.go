package response

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/go-chi/chi"
	"github.com/rianekacahya/errors"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	e := chi.NewRouter()
	e.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		err := errors.New(errors.BADREQUEST, errors.Message("Error"))
		Error(w, err)
	})

	ts := httptest.NewServer(e)
	defer ts.Close()
	var cl http.Client
	req, _ := http.NewRequest("GET", ts.URL+"/ping", nil)
	resp, _ := cl.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, `{"message":"Error"}`, string(body))
}

func TestRender(t *testing.T) {
	e := chi.NewRouter()
	e.Get("/pong", func(w http.ResponseWriter, r *http.Request) {
		Render(w, http.StatusOK, nil)
	})

	ts := httptest.NewServer(e)
	defer ts.Close()
	var cl http.Client
	req, _ := http.NewRequest("GET", ts.URL+"/pong", nil)
	resp, _ := cl.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, `{"message":"success"}`, string(body))
}
