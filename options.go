package jpush

type Options struct {
	ThirdPartyChannel map[string]inteface{} `json:"third_party_channel,omitempty"` // 生产环境为 true
}
