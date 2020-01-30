package models

import (
	"encoding/json"
	"io"
	"net/http"
)

//App error codes

type ApiResponse struct {
	Status        string            `json:"status"`                  // Message to be display to the end user without debugging information
	StatusCode    int               `json:"statusCode"`              // The http status code
	Message       string            `json:"message,omitempty"`       // The http status code
	DetailedError string            `json:"detailedError,omitempty"` // Internal error string to help the developer
	ErrorCode     string            `json:"errorCode,omitempty"`     // Internal error string to help the developer
	Data          interface{}       `json:"data,omitempty"`
	Headers       map[string]string `json:"_,omitempty"`
	Cookies       *http.Cookie      `json:"_,omitempty"`
}

func NewFRes(statusCode int, message string, detailedError string, errorCode string) *ApiResponse {
	fr := &ApiResponse{}
	fr.Status = FAILED_STATUS
	fr.StatusCode = statusCode
	fr.Message = message
	fr.DetailedError = detailedError
	fr.ErrorCode = errorCode
	return fr
}

func NewSRes(statusCode int, message string, data interface{}) *ApiResponse {
	sr := &ApiResponse{}
	sr.Status = SUCCESS_STATUS
	sr.StatusCode = statusCode
	sr.Message = message
	sr.Data = data

	return sr
}

func (ar *ApiResponse) ToJson() string {
	b, err := json.Marshal(ar)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func ApiResponseFromJson(data io.Reader) *ApiResponse {
	var apiResponse *ApiResponse
	json.NewDecoder(data).Decode(&apiResponse)
	return apiResponse
}
