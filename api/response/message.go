package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
	Code    uint   `json:"code"`
}

func MessageResponse(w http.ResponseWriter, message string, code uint) {
	response := Message{Message: message, Code: code}
	encodedMessage, _ := json.Marshal(response)

	w.Header().Add("content-type", "application/json")
	w.Header().Add("charset", "utf-8")
	fmt.Fprint(w, string(encodedMessage))
}
