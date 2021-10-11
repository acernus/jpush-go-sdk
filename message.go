package jpush

type Message struct {
	MsgContent  string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}
