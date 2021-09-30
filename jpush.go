package jpush

import "sync"

type Client struct {
}

var clientInstance *Client
var once = sync.Once{}

func NewClient() *Client {
	once.Do(func() {
		clientInstance = &Client{}
	})
	return clientInstance
}

func (s *Client) Push() error {
	return nil
}
