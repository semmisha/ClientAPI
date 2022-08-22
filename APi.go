package ClientAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func NewConnection(channel string) *Message {
	return &Message{Channel: channel}
}

func (m Message) Write(p []byte) (n int, err error) {

	m.Text = string(p)
	marshMellow, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)

	}
	client := http.Client{}

	resp, err := client.Post("http://127.0.0.1:9999/PostMessage", "application/json ", bytes.NewBuffer(marshMellow))
	if err != nil {

		fmt.Println("error")

	}

	fmt.Println(resp.StatusCode)
	resdpBody, _ := io.ReadAll(resp.Body)

	fmt.Println(string(resdpBody))

	return 0, err
}
