package client

import (
	"fmt"
	"io"
	"net/http"
)

func DoRequest(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s | body: %s", resp.Status, string(body))
	}

	return body, nil
}
