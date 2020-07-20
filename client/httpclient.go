package client

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {

	c := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, errors.New("Error in httpclient ")
	}

	resp, cerr := c.Do(req)

	if cerr != nil {
		return nil, errors.New("Error in httpclient ")
	}

	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes, nil

}
