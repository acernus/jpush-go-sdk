package jpush

type CallbackType int

const (
	CallbackTypeDelivery            CallbackType = 1
	CallbackTypeClick               CallbackType = 2
	CallbackTypeDeliveryAndClick    CallbackType = 3
	CallbackTypePushed              CallbackType = 8
	CallbackTypePushedAndDelivery   CallbackType = 9
	CallbackTypePushedAndClick      CallbackType = 10
	CallbackTypePushedDeliveryClick CallbackType = 11
)

type Callback struct {
	Url    string                 `json:"url"`
	Params map[string]interface{} `json:"params"`
	Type   CallbackType           `json:"type"`
}
