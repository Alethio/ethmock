package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Client represents a http rpc client
type Client struct {
	// Client represents a http rpc client
	client *http.Client
	url    string
}

// Request makes the call to the client
func (c *Client) Request(jsonRequest []byte) ([]byte, error) {
	httpRequest, err := http.NewRequest("POST", c.url, bytes.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}
	defer httpRequest.Body.Close()

	httpRequest.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// New initializes a Client and returns it
func New(url string) (*Client, error) {
	var httpClient = &http.Client{Transport: &http.Transport{}}

	p := &Client{
		url:    url,
		client: httpClient,
	}

	return p, nil
}
