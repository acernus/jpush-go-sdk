package jpush

type ApnsProduction bool

const (
	Production ApnsProduction = true
	Test       ApnsProduction = false
)

type Options struct {
	ApnsProduction ApnsProduction `json:"apns_production"`
}
