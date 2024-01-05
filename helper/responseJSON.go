package helper

// ResponseJSON will be used by us to send standard response for all the api
type ResponseJSON struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
