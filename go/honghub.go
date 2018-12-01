package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// MessageFromLine Core
type MessageFromLine struct {
	Event       []Event `json:"events"`
	Destination string  `json:"destination"`
}

// Event for MessageFromLine
type Event struct {
	Type       string  `json:"type"`
	ReplyToken string  `json:"replyToken"`
	Source     Source  `json:"source"`
	Timestamp  float64 `json:"timestamp"`
	Message    Message `json:"message"`
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

type Reply struct {
	ReplyToken string            `json:"replyToken"`
	Messages   []MessageForReply `json:"messages"`
}
type MessageForReply struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// MessageReceive For Webhook
func MessageReceive(w http.ResponseWriter, r *http.Request) {
	// dec := json.NewDecoder(strings.NewReader(jsonStream))
	// fmt.Fprintf(w, "Hello World! is updated at 22:37")
	decoder := json.NewDecoder(r.Body)
	var rs MessageFromLine
	err := decoder.Decode(&rs)
	if err != nil {
		fmt.Println("Tst: ", err)
		JSONResponse(w, http.StatusOK, rs)
		return
	}

	// fmt.Println("raw body", string(byt))
	ReplyMessage(rs, r)
	JSONResponse(w, http.StatusOK, rs)
}

// ReplyMessage For reply message
func ReplyMessage(rs MessageFromLine, r *http.Request) {
	// url := "https://api.line.me/v2/bot/message/reply"
	// client := &http.Client{}
	// var test Reply
	// var text MessageForReply
	// test.ReplyToken = rs.Event[0].ReplyToken
	// text.Type = "text"
	// text.Text = "นาย หล่อมาก"
	// test.Messages = append(test.Messages, text)
	// bytesRepresentation, err := json.Marshal(test)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", "Bearer {t38SlKJyLMd1kepY/2JS/UTB32hiNtR1lxF/SModWBDB4Xn9+3ofh0Gu/iKrn1sEZt02dwvyTjlapNJttYV+NbBMLgcHCwjdTyNweSXdulnQRwzxWA/VKRsfPiSqaXJSdQCCXpUfqVGi1RsMC4QhvwdB04t89/1O/w1cDnyilFU=}")
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(resp)
	// defer resp.Body.Close()

	bot, err := linebot.New("cbf2f7378ec06aa02731da2a6182ff90", "t38SlKJLMd1kepY/2JS/UTB32hiNtR1lxF/SModWBDB4Xn9+3ofh0Gu/iKrn1sEZt02dwvyTjlapNJttYV+NbBMLgcHCwjdTyNweSXdulnQRwzxWA/VKRsfPiSqaXJSdQCCXpUfqVGi1RsMC4QhvwdB04t89/1O/w1cDnyilFU=")
	if err != nil {
		fmt.Println("1: ", err)
		// Do something when something bad happened.
	}
	// events, err := bot.ParseRequest(r)
	// if err != nil {
	// 	fmt.Println("2: ", err)
	// 	// Do something when something bad happened.
	// }
	for _, event := range rs.Event {
		if event.Type == "text" {
			// var messages []linebot.Message

			// append some message to messages
			leftBtn := linebot.NewMessageAction("left", "left clicked")
			rightBtn := linebot.NewMessageAction("right", "right clicked")
			template := linebot.NewConfirmTemplate("Hello World", leftBtn, rightBtn)

			message := linebot.NewTemplateMessage("นาย หล่อมาก ", template)
			_, err := bot.ReplyMessage(event.ReplyToken, message).Do()
			if err != nil {
				// Do something when some bad happened
				fmt.Println("DDDDDD2: ", err)

			}
		}
	}

	// fmt.Println("Test: ", events)
	// fmt.Println("Test: ", rs)

}
