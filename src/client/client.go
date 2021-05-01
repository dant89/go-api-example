package client

import (
	"io"
	"net/http"
)

type Client struct {
	host string
}

func NewClient(host string) Client {
	return Client{host}
}

func (c Client) Get(uri string) string {
	resp, err := http.Get(c.host + "/" + uri)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}

func (c Client) GetUsers() string {
	return c.Get("users")
}

func (c Client) GetUser(userId string) string {
	return c.Get("users/" + userId)
}
