package jpush

type ApnsProduction bool

const (
	Production ApnsProduction = true
	Test       ApnsProduction = false
)

const (
	ThirdPartyChannelHuawei = "huawei"
	ThirdPartyChannelXiaomi = "xiaomi"
	ThirdPartyChannelMeizu  = "meizu"
	ThirdPartyChannelOppo   = "oppo"
	ThirdPartyChannelVivo   = "vivo"
	ThirdPartyChannelAsus   = "asus"

	DistributionJpush         = "jpush"
	DistributionOspush        = "ospush"
	DistributionSecondaryPush = "secondary_push"
	DistributionFirstOspush   = "first_ospush"
)

type Options struct {
	TimeToLive        *int                          `json:"time_to_live,omitempty"`
	ThirdPartyChannel map[string]*ThirdPartyChannel `json:"third_party_channel,omitempty"`
	ApnsProduction    ApnsProduction                `json:"apns_production"`
}

type ThirdPartyChannel struct {
	Distribution          string                 `json:"distribution,omitempty"`
	DistributionFcm       string                 `json:"distribution_fcm,omitempty"`
	DistributionCustomize string                 `json:"distribution_customize,omitempty"`
	ChannelID             string                 `json:"channel_id,omitempty"`
	SkipQuota             bool                   `json:"skip_quota,omitempty"`
	Classification        int                    `json:"classification,omitempty"`
	PushMode              int                    `json:"push_mode,omitempty"`
	Importance            string                 `json:"importance,omitempty"`
	Urgency               string                 `json:"urgency,omitempty"`
	Category              string                 `json:"category,omitempty"`
	LargeIcon             string                 `json:"large_icon,omitempty"`
	SmallIconUri          string                 `json:"small_icon_uri,omitempty"`
	SmallIconColor        string                 `json:"small_icon_color,omitempty"`
	Style                 int                    `json:"style,omitempty"`
	BigText               string                 `json:"big_text,omitempty"`
	Inbox                 map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath            string                 `json:"big_pic_path,omitempty"`
	OnlyUseVendorStyle    bool                   `json:"only_use_vendor_style,omitempty"`
}

func NewOptions() *Options {
	return &Options{}
}

func (opt *Options) FillThirdPartyChannel(channelName string, param *ThirdPartyChannel) *Options {
	if opt.ThirdPartyChannel == nil {
		opt.ThirdPartyChannel = make(map[string]*ThirdPartyChannel)
	}
	opt.ThirdPartyChannel[channelName] = param
	return opt
}

func (opt *Options) FillTimeToLive(timeToLive int) *Options {
	opt.TimeToLive = &timeToLive
	return opt
}
