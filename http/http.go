package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

func MakeHttpCall(url string) (*Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &Response{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil
}
