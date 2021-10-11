package jpush

import "encoding/json"

type PlatformType string

const (
	PlatformAll      PlatformType = "all"
	PlatformiOS      PlatformType = "ios"
	PlatformAndroid  PlatformType = "android"
	PlatformQuickApp PlatformType = "quickapp"
)

type Platform struct {
	All    bool
	Values []PlatformType
}

func (p *Platform) MarshalJSON() ([]byte, error) {
	if p.All || p.containsAll() {
		return []byte(PlatformAll), nil
	}
	return json.Marshal(p.Values)
}

func (p *Platform) containsAll() bool {
	for _, item := range p.Values {
		if item == PlatformAll {
			return true
		}
	}
	return false
}

func NewPlatform() *Platform {
	return &Platform{
		Values: make([]PlatformType, 0),
	}
}

func (p *Platform) UseAll() *Platform {
	p.All = true
	return p
}

func (p *Platform) UsePlatform(platform ...PlatformType) *Platform {
	for _, item := range platform {
		p.Values = append(p.Values, item)
	}
	return p
}
