package helper

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url, msg string, headers map[string]string) (string, error) {

	var err error
	var req *http.Request
	var resp *http.Response
	var body []byte

	rd := strings.NewReader(msg)
	if req, err = http.NewRequest("POST", url, rd); err != nil {
		return "", err
	}

	for key, header := range headers {
		req.Header.Set(key, header)
	}

	client := &http.Client{}
	if resp, err = client.Do(req); err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(body), err

}
