package tracesmith

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	Endpoint string
	APIKey   string
}

func NewClient() *Client {
	return &Client{
		Endpoint: os.Getenv("LANGCHAIN_ENDPOINT"),
		APIKey:   os.Getenv("LANGCHAIN_API_KEY"),
	}
}

func (c *Client) sendRequest(method, path string, data map[string]interface{}) error {
	if c.Endpoint == "" || c.APIKey == "" {
		return fmt.Errorf("missing environment variables: LANGCHAIN_ENDPOINT or LANGCHAIN_API_KEY")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	log.Debugf("Payload being sent: %s", string(jsonData))

	req, err := http.NewRequest(method, c.Endpoint+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error response: %s", resp.Status)
	}
	return nil
}
