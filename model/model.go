// Package model holds the object for the request and response payload
package model

const (
	// LogStrRequest log string key
	LogStrRequest = "request"
	// LogStrResponse log string key
	LogStrResponse = "response"
	// BaseURL for Falcon-X
	BaseURL = "https://api.column.com"
	// APIKey for authentication
	APIKey = ""
	// HeaderContentTypeKey content type header
	HeaderContentTypeKey = "Content-Type"
	// HeaderContentTypeValue content type header
	HeaderContentTypeValue = "application/json"
	// HeaderContentAcceptKey content accept header
	HeaderContentAcceptKey = "Accept"
	// ProxyIPAddressAndPort proxy ip address and port value
	ProxyIPAddressAndPort = "http://3.22.106.10:8888" //"http://ec2-3-22-106-10.us-east-2.compute.amazonaws.com:8888"

	// IdempotencyKey header key.
	IdempotencyKey = "Idempotency-Key"
)

type (
	// Response schema for all endpoints
	Response struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	// ErrorResponse schema for error responses
	ErrorResponse struct {
		Type             string      `json:"type"`
		Code             string      `json:"code"`
		Message          string      `json:"message"`
		DocumentationURL string      `json:"documentation_url"`
		Details          interface{} `json:"details"`
	}

	// ColumnErrorCode string representation of an error code on Column
	ColumnErrorCode string
)
