package response

// Describes an outgoing response structure
type Response struct {

	// Business payload
	Data interface{} `json:"data,omitempty"`
}
