package json

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	HttpClient http.Client
}

func (c *Client) Post(url string, body interface{}, resp interface{}) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return errors.WithStack(err)
	}

	r, err := http.NewRequest("POST", url, bytes.NewReader(jsonBody))
	if err != nil {
		return errors.WithStack(err)
	}

	jsonResponse, err := c.HttpClient.Do(r)
	if err != nil {
		return errors.WithStack(err)
	}
	defer jsonResponse.Body.Close()

	responseBytes, err := ioutil.ReadAll(jsonResponse.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := json.Unmarshal(responseBytes, resp); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
