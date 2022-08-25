package ClientAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	URL     string `json:"url"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func NewConnection(channel string, url string) *Message {
	return &Message{
		Channel: channel,
		URL:     url,
	}
}

func (m *Message) Write(p []byte) (n int, err error) {

	m.Text = bytes.NewBuffer(p).String()

	marshMellow, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)

	}
	client := http.Client{}

	resp, err := client.Post(m.URL, "application/json ", bytes.NewBuffer(marshMellow))
	if err != nil {

		fmt.Println("error")

	}

	fmt.Println(resp.StatusCode)
	resdpBody, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)

	}

	fmt.Println(string(resdpBody), "")

	return 0, err

}
