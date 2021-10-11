package jpush

type PushReq struct {
	Platform     *Platform     `json:"platform"`
	Audience     *Audience     `json:"audience"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	Callback     *Callback     `json:"callback,omitempty"`
	CID          string        `json:"cid,omitempty"`
}

type PushResp struct {
	SendNo string `json:"send_no"`
	MsgID  string `json:"msg_id"`
}

type CidType string

const (
	CidTypePush     CidType = "push"
	CidTypeSchedule CidType = "schedule"
)

type CidReq struct {
	Count int     `json:"count,omitempty"`
	Type  CidType `json:"type,omitempty"`
}

type CidResp struct {
	CidList []string `json:"cidlist"`
}