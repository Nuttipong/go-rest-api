package viewmodels

// HttpResponseVM to represent model to client
type HttpResponseVM struct {
	StatusCode int    `json:"statusCode"`
	Result     Result `json:"result,omitempty"`
}

// Result to generic
type Result interface{}
