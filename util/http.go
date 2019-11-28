package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GET http get
func GET(path string) (resp *http.Response, bs []byte, err error) {
	resp, err = http.Get(path)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("http get error, status = %d", resp.StatusCode)
		return
	}
	bs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// DoPostXml
func DoPostXml(url string, postData string) ([]byte, error) {
	resp, err := http.Post(url, "application/xml; charset=utf-8", strings.NewReader(postData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
