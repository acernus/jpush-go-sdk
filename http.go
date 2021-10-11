package jpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpPostJSON(reqURL string, params interface{}, headers map[string]string) (string, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", reqURL, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("resp code not ok, body: %s", string(body))
	}
	return string(body), nil
}

func HttpPostForm(reqURL string, params map[string]interface{}, headers map[string]string) (string, error) {
	p := url.Values{}
	for k, v := range params {
		data, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		p.Set(k, string(data))
	}

	req, err := http.NewRequest("POST", reqURL, strings.NewReader(p.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("resp code not ok, body: %s", string(body))
	}

	return string(body), nil
}

func HttpGet(reqURL string, params map[string]string, headers map[string]string) (string, error) {
	p := url.Values{}
	for k, v := range params {
		if v == "" {
			continue
		}
		p.Set(k, v)
	}

	targetURL := reqURL + "?" + p.Encode()
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("resp code not ok, body: %s", string(body))
	}
	return string(body), nil
}
