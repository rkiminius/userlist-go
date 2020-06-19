package userlist

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status int      `json:"status,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

type Client struct{}

func (c *Client) Get(endpoint string) {
	c.Request("GET", endpoint, nil)
}

func (c *Client) Post(endpoint string, payload interface{}) {
	c.Request("POST", endpoint, payload)
}

func (c *Client) Request(method, endpoint string, payload interface{}) error {
	client := &http.Client{}

	body, _ := json.Marshal(payload)

	req, err := http.NewRequest(method, PUSH_ENDPOINT+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Push "+PUSH_KEY)

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var res Response
		err = json.NewDecoder(response.Body).Decode(&res)
		if err != nil {
			log.Println("error ", err)
		}
		log.Println(res)
	}

	return nil
}

func (c *Client) Delete(endpoint string) {
	c.Request("DELETE", endpoint, nil)
}
