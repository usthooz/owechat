package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
