package jpush

type ApnsProduction string

const (
	Production ApnsProduction = "True"
	Test       ApnsProduction = "False"
)

type Options struct {
	ApnsProduction string `json:"apns_production,omitempty"`
}
