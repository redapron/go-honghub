package swagger

import (
	json "encoding/json"
	"fmt"
	"net/http"
	"time"
)

// MessageFromLine Core
type MessageFromLine struct {
	Event       Event  `json:"events"`
	Destination string `json:"destination"`
}

// Event for MessageFromLine
type Event struct {
	Type       string     `json:"type"`
	ReplyToken string     `json:"replyToken"`
	Source     Source     `json:"source"`
	Timestamp  *time.Time `json:"timestamp"`
	Message    Message    `json:"message"`
}

// Message for Event
type Message struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

// Source for Event
type Source struct {
	RoomID string `json:"roomId"`
	UserID string `json:"userId"`
	Type   string `json:"type"`
}

// MessageReceive For Webhook
func MessageReceive(w http.ResponseWriter, r *http.Request) {
	// dec := json.NewDecoder(strings.NewReader(jsonStream))
	// fmt.Fprintf(w, "Hello World! is updated at 22:37")
	decoder := json.NewDecoder(r.Body)
	var rs MessageFromLine
	err := decoder.Decode(&rs)
	if err != nil {
		JSONResponse(w, http.StatusOK, nil)
	}

	fmt.Println("raw body", string(byt))
	ReplyMessage(rs)
	JSONResponse(w, http.StatusOK, nil)
}

// ReplyMessage For reply message
func ReplyMessage(rs MessageFromLine) {
	fmt.Println("Test : ", rs)
}
