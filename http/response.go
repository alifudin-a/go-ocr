package http

// Response : custom response struct
type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message,omitmepty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}
