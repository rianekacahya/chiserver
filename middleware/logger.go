package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
	"encoding/json"
	"github.com/rianekacahya/logger"
	"go.uber.org/zap"
)

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var(
			request []byte
			response []byte
		)

		if r.Body != nil {
			request, _ = ioutil.ReadAll(r.Body)
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(request))
		recorder := httptest.NewRecorder()

		next.ServeHTTP(recorder, r)

		// Capture all header
		for key := range recorder.Header() {
			w.Header().Set(key, recorder.Header().Get(key))
		}

		response, _ = ioutil.ReadAll(recorder.Body)

		if recorder.Code >= http.StatusOK && recorder.Code < http.StatusMultipleChoices {
			reqMessage := json.RawMessage(`""`)
			resMessage := json.RawMessage(`""`)

			if len(request) > 0 {
				reqMessage = json.RawMessage(request)
			}

			if len(response) > 0 {
				resMessage = json.RawMessage(response)
			}

			logger.Info(
				"http",
				zap.Int("status", recorder.Code),
				zap.String("time", time.Now().Format(time.RFC1123Z)),
				zap.String("hostname", r.Host),
				zap.String("user_agent", r.UserAgent()),
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("query", r.URL.RawQuery),
				zap.Any("req", reqMessage),
				zap.Any("res", resMessage),
			)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(recorder.Code)
			w.Write(response)
		}
	}

	return http.HandlerFunc(fn)
}
