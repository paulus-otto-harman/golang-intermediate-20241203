package domain

type HTTPResponse struct {
	Status      bool        `json:"status"`
	ErrorCode   string      `json:"error_code,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
