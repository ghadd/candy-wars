package api

// Struct that symbolizes the query, which is being sent along with inline button press
type CallBackQuery struct {
	ID           string        `json:"id"`
	FromUser     User          `json:"from"`
	Message      UpdateMessage `json:"message"`
	CallBackData string        `json:"data"`
}
