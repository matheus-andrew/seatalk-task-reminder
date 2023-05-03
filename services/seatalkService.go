package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type MessageBot struct {
	Tag      string   `json:"tag"`
	Markdown Markdown `json:"markdown"`
}

type Markdown struct {
	Content            string        `json:"content"`
	MentionedList      []interface{} `json:"mentioned_list"`
	MentionedEmailList []interface{} `json:"mentioned_email_list"`
	AtAll              bool          `json:"at_all"`
}

func SendMessage(content string) {
	seatalk_group_id := os.Getenv("SEATALK_GROUP_ID")
	url := "https://openapi.seatalk.io/webhook/group/" + seatalk_group_id

	bot := &MessageBot{
		Tag: "markdown",
		Markdown: Markdown{
			Content:            content,
			MentionedList:      nil,
			MentionedEmailList: nil,
			AtAll:              false,
		},
	}

	result, _ := json.Marshal(bot)

	request, error := http.NewRequest("POST", url, bytes.NewBuffer(result))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
}
