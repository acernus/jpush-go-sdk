package model

type PlatformType string

const (
	PlatformAll      PlatformType = "all"
	PlatformiOS      PlatformType = "ios"
	PlatformAndroid  PlatformType = "android"
	PlatformQuickApp PlatformType = "quickapp"
)

type PushReq struct {
	CID          string        `json:"cid"`
	Platform     PlatformType  `json:"platform"`
	Audience     *Audience     `json:"audience"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
}
