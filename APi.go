package ClientAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TLGClient struct {
	URL     string `json:"url"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func NewConnection(channel string, url string) *TLGClient {
	return &TLGClient{
		Channel: channel,
		URL:     url,
	}
}

func (m *TLGClient) Write(p []byte) (n int, err error) {

	m.Text = bytes.NewBuffer(p).String()

	marshMellow, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)

	}
	client := http.Client{}

	resp, err := client.Post(m.URL, "application/json ", bytes.NewBuffer(marshMellow))
	if err != nil {

		fmt.Println("Post request error: ", err)

	}

	fmt.Println("Reply status code: ", resp.StatusCode)

	return 0, err

}
