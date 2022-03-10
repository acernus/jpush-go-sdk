package jpush

type AndroidIntent struct {
	Url string `json:"url"`
}

type NotificationAndroid struct {
	Alert             string                 `json:"alert"`
	Title             string                 `json:"title,omitempty"`
	BuilderId         int                    `json:"builder_id,omitempty"`
	LargeIcon         string                 `json:"large_icon,omitempty"`
	Intent            *AndroidIntent         `json:"intent,omitempty"`
	DisplayForeground string                 `json:"display_foreground,omitempty"`
	Priority          int                    `json:"priority,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
}

type IosPayload struct {
	Title    string `json:"title,omitempty"`
	SubTitle string `json:"sub_title,omitempty"`
	Body     string `json:"body,omitempty"`
}

type NotificationIOS struct {
	Alert            *IosPayload            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            string                 `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	ThreadId         string                 `json:"thread-id,omitempty"`
}

type VOIP struct {
	Key string `json:"key"`
}

type QuickApp struct {
	Alert  string                 `json:"alert"`
	Title  string                 `json:"title"`
	Page   string                 `json:"page"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}

type Notification struct {
	NotificationAndroid *NotificationAndroid   `json:"android,omitempty"`
	NotificationIOS     *NotificationIOS       `json:"ios,omitempty"`
	InappMessage        *InappMessage          `json:"inapp_message,omitempty"`
	Voip                map[string]interface{} `json:"voip,omitempty"`
	Quickapp            *QuickApp              `json:"quickapp,omitempty"`
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) IOS(req *NotificationIOS) *Notification {
	n.NotificationIOS = req
	return n
}

func (n *Notification) Android(req *NotificationAndroid) *Notification {
	n.NotificationAndroid = req
	return n
}

func (n *Notification) UseInAppMessage() *Notification {
	n.InappMessage = &InappMessage{InappMessage: true}
	return n
}
func (n *Notification) VOIP(req map[string]interface{}) *Notification {
	n.Voip = req
	return n
}
