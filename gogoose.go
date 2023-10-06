package gogoose

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GooseClient struct {
	url string
}

func New(localHostUrl string) *GooseClient {
	return &GooseClient{url: localHostUrl}
}

func (gc *GooseClient) All() (*http.Response, error) {
	resp, err := http.Get(gc.url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (gc *GooseClient) Send(jsonData interface{}) (*http.Response, error) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(gc.url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (gc *GooseClient) Update(jsonData interface{}) (*http.Response, error) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(gc.url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (gc *GooseClient) Delete(jsonData interface{}) (*http.Response, error) {
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(gc.url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
