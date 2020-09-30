package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

/*
ConnectUsingHTTP define
*/
func ConnectUsingHTTP(method, url string, body []byte, header map[string][]string) (code int, contentType string, buff []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return 1, "", nil, err
	}
	for k, h := range header {
		for _, v := range h {
			req.Header.Add(k, v)
		}
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 2, "", nil, err
	}
	defer resp.Body.Close()
	buff, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 3, "", nil, err
	}
	return resp.StatusCode, resp.Header.Get("Content-Type"), buff, nil
}
