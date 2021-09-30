package model

type NotificationAndroid struct {
	Alert     string `json:"alert"`
	Title     string `json:"title"`
	BuilderId int    `json:"builder_id"`
	LargeIcon string `json:"large_icon"`
	Intent    struct {
		Url string `json:"url"`
	} `json:"intent"`
	Extras interface{} `json:"extras"`
}

type NotificationIOS struct {
	Alert    string      `json:"alert"`
	Sound    string      `json:"sound"`
	Badge    string      `json:"badge"`
	ThreadId string      `json:"thread-id"`
	Extras   interface{} `json:"extras"`
}

type VOIP struct {
	Key string `json:"key"`
}

type QuickApp struct {
	Alert string `json:"alert"`
	Title string `json:"title"`
	Page  string `json:"page"`
}

type Notification struct {
	Android  NotificationAndroid `json:"android,omitempty"`
	IOS      NotificationIOS     `json:"ios,omitempty"`
	Voip     *VOIP               `json:"voip,omitempty"`
	Quickapp *QuickApp           `json:"quickapp,omitempty"`
}
