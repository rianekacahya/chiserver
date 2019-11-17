package response

import (
	"encoding/json"
	"github.com/rianekacahya/errors"
	"net/http"
)

type response struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Error(c http.ResponseWriter, err error) {
	var(
		status int
		response = new(response)
	)

	if err != nil {
		// Mapping Status
		errorStatus := errors.GetStatus(err)
		errorMessage := errors.GetError(err)
		switch errorStatus {
		case errors.GENERIC:
			status = http.StatusInternalServerError
		case errors.FORBIDDEN:
			status = http.StatusForbidden
		case errors.BADREQUEST:
			status = http.StatusBadRequest
		case errors.NOTFOUND:
			status = http.StatusNotFound
		case errors.UNAUTHORIZED:
			status = http.StatusUnauthorized
		default:
			status = http.StatusInternalServerError
			response.Message = err.Error()
		}

		if errorStatus != errors.NOTYPE {
			switch errorMessage.(type) {
			case errors.Custom:
				response.Message = errorMessage.Error()
			default:
				response.Message = errorMessage
			}
		}
	}

	res, _ := json.Marshal(response)
	c.Header().Set("Content-Type", "application/json")
	c.WriteHeader(status)
	c.Write(res)
}

func Render(c http.ResponseWriter, status int, data interface{}) {
	var response = new(response)

	response.Message = "success"
	response.Data = data

	res, _ := json.Marshal(response)
	c.Header().Set("Content-Type", "application/json")
	c.WriteHeader(status)
	c.Write(res)
}
