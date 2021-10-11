package jpush

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
)

type Client struct {
	Authorization string
}

func NewClient(appKey, masterSecret string) *Client {
	return &Client{
		Authorization: base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", appKey, masterSecret))),
	}
}

func (s *Client) Push(req *PushReq) (string, error) {
	resp, err := HttpPostJSON(PushApi, req, map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", s.Authorization),
	})
	if err != nil {
		return "", err
	}
	var pushResp PushResp
	err = json.Unmarshal([]byte(resp), &pushResp)
	if err != nil {
		return "", err
	}

	return pushResp.MsgID, nil
}

func (s *Client) PushValidate(req *PushReq) (string, error) {
	resp, err := HttpPostJSON(PushValidateApi, req, map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", s.Authorization),
	})
	if err != nil {
		return "", err
	}
	var pushResp PushResp
	err = json.Unmarshal([]byte(resp), &pushResp)
	if err != nil {
		return "", err
	}

	return pushResp.MsgID, nil
}

func (s *Client) GetCID(req *CidReq) ([]string, error) {
	resp, err := HttpGet(CIDApi, map[string]string{
		"count": strconv.Itoa(req.Count),
		"type":  string(req.Type),
	}, map[string]string{
		"Authorization": fmt.Sprintf("Basic %s", s.Authorization),
	})
	if err != nil {
		return nil, err
	}
	var cidResp CidResp
	err = json.Unmarshal([]byte(resp), &cidResp)
	if err != nil {
		return nil, err
	}

	return cidResp.CidList, nil
}
