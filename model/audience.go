package model

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
	*AudienceTag
	*AudienceAlias
	*AudienceRegistration
}
