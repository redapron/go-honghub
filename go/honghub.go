package swagger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

	bot, err := linebot.New("cbf2f7378ec06aa02731da2a6182ff90", "oZZxTbxzU/hG2dAL90WeofEkb9iXy7/19qvWzXWpBSORORk8ZCTGRMoLWd6PsqNyZt02dwvyTjlapNJttYV+NbBMLgcHCwjdTyNweSXdullp2w4/V061Vqs70ZVd4OfGFD9B9kx5RH1ntZh4zNYoZwdB04t89/1O/w1cDnyilFU=")
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
		fmt.Println("Test: ", event.Type)
		if event.Message.Type == "text" {
			var message linebot.SendingMessage
			if strings.Index(event.Message.Text, "1") >= 0 {
				leftBtn := linebot.NewMessageAction("จองห้องหน่อยสิ", "จองห้อง")
				rightBtn := linebot.NewMessageAction("มาทักทาย", "ทักทาย")
				template := linebot.NewConfirmTemplate("สวัสดีคุณ ... , นุ่นคือบอทที่จะช่วยให้คุณจองห้องประชุมได้ง่ายๆ ^^", leftBtn, rightBtn)
				message = linebot.NewTemplateMessage("", template)
			} else if strings.Index(event.Message.Text, "2") >= 0 {
				message = linebot.NewTextMessage("ไม่ทราบว่าคุณ...ต้องการประชุมที่ตึกไหนคะ ").
					WithQuickReplies(linebot.NewQuickReplyItems(
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("ตึกแจ้ง 1", "ตึกแจ้ง 1")),
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("KBTG", "KBTG")),
					))

			} else if strings.Index(event.Message.Text, "3") >= 0 || strings.Index(event.Message.Text, "3") >= 0 {
				message = linebot.NewTextMessage("กรุณาพิมพ์หมายเลขเพื่อเลือกประเภทห้องประชุมค่ะ ").
					WithQuickReplies(linebot.NewQuickReplyItems(
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("ห้อง Whiteboard", "ห้อง Whiteboard")),
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("ห้อง Video conference", "ห้อง Video conference")),
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("ห้อง Project", "ห้อง Project")),
					))
			} else if strings.Index(event.Message.Text, "4") >= 0 {
				message = linebot.NewTextMessage("ไม่ทราบว่าคุณ...ต้องการจองห้องไหนคะ").
					WithQuickReplies(linebot.NewQuickReplyItems(
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("ห้อง 1", "ห้อง 1")),
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("ห้อง 2", "ห้อง 2")),
						linebot.NewQuickReplyButton(
							"",
							linebot.NewMessageAction("ห้อง 3", "ห้อง 3")),
					))

			} else if strings.Index(event.Message.Text, "5") >= 0 {
				template := linebot.NewButtonsTemplate(
					"", "", "ช่วงเวลาไหนคะ",
					linebot.NewDatetimePickerAction("date", "DATE", "date", "", "", ""),
					linebot.NewDatetimePickerAction("time", "TIME", "time", "", "", ""),
					linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
				)
				message = linebot.NewTemplateMessage("", template)

			} else {
				message = linebot.NewTextMessage("นุ่นต้องขอโทษด้วยค่ะ นุ่นไม่เข้าใจ")
			}
			_, err := bot.ReplyMessage(event.ReplyToken, message).Do()
			if err != nil {
				fmt.Println("DDDDDD2: ", err)
			}
		}
	}

	// fmt.Println("Test: ", events)
	// fmt.Println("Test: ", rs)

}
