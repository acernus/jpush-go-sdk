package jpush

import "encoding/json"

type AudienceValue struct {
	*AudienceTag
	*AudienceAlias
	*AudienceRegistration
}
type AudienceTag struct {
	Tag    []string `json:"tag,omitempty"`
	TagAnd []string `json:"tag_and,omitempty"`
	TagNot []string `json:"tag_not,omitempty"`
}

type AudienceAlias struct {
	Alias []string `json:"alias,omitempty"`
}
type AudienceRegistration struct {
	RegistrationID []string `json:"registration_id,omitempty"`
}

type Audience struct {
	All   bool
	Value *AudienceValue
}

func (a *Audience) MarshalJSON() ([]byte, error) {
	if a.All {
		return []byte("all"), nil
	}
	return json.Marshal(a.Value)
}

func NewAudience() *Audience {
	return &Audience{
		All: false,
		Value: &AudienceValue{
			&AudienceTag{},
			&AudienceAlias{},
			&AudienceRegistration{},
		},
	}
}

func (a *Audience) UseAll() *Audience {
	a.All = true
	return a
}

func (a *Audience) TagListOr(tags ...string) *Audience {
	if a.Value.Tag == nil {
		a.Value.Tag = make([]string, 0)
	}
	for _, tag := range tags {
		a.Value.Tag = append(a.Value.Tag, tag)
	}
	return a
}

func (a *Audience) TagListAnd(tags ...string) *Audience {
	if a.Value.TagAnd == nil {
		a.Value.TagAnd = make([]string, 0)
	}
	for _, tag := range tags {
		a.Value.TagAnd = append(a.Value.TagAnd, tag)
	}
	return a
}

func (a *Audience) TagListNot(tags ...string) *Audience {
	if a.Value.TagNot == nil {
		a.Value.TagNot = make([]string, 0)
	}
	for _, tag := range tags {
		a.Value.TagNot = append(a.Value.TagNot, tag)
	}
	return a
}

func (a *Audience) AliasListOr(alias ...string) *Audience {
	if a.Value.Alias == nil {
		a.Value.Alias = make([]string, 0)
	}
	for _, item := range alias {
		a.Value.Alias = append(a.Value.Alias, item)
	}
	return a
}

func (a *Audience) RegistrationList(registrations ...string) *Audience {
	if a.Value.RegistrationID == nil {
		a.Value.RegistrationID = make([]string, 0)
	}
	for _, registration := range registrations {
		a.Value.RegistrationID = append(a.Value.RegistrationID, registration)
	}
	return a
}
