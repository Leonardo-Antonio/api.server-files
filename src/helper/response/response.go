package response

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Error       bool        `json:"error"`
	Data        interface{} `json:"data"`
}

func Successful(message string, data interface{}) *response {
	return &response{
		MessageType: "message",
		Message:     message,
		Error:       false,
		Data:        data,
	}
}

func Err(message string, data interface{}) *response {
	return &response{
		MessageType: "error",
		Message:     message,
		Error:       true,
		Data:        data,
	}
}
