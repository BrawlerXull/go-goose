package gogoose

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GooseClient struct {
	Url string
}

func New(localHostUrl string) *GooseClient {
	return &GooseClient{Url: localHostUrl}
}

func (gc *GooseClient) All() (*http.Response, error) {
	resp, err := http.Get(gc.Url + "all")
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

	resp, err := http.Post(gc.Url+"send", "application/json", bytes.NewBuffer(jsonBytes))
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

	resp, err := http.Post(gc.Url+"update", "application/json", bytes.NewBuffer(jsonBytes))
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

	resp, err := http.Post(gc.Url+"delete", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
