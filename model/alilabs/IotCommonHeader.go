package alilabs

// IotCommonHeader 结构体
type IotCommonHeader struct {
	// payLoadVersion
	PayLoadVersion int64 `json:"pay_load_version,omitempty" xml:"pay_load_version,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// messageId，区分请求使用
	MessageId string `json:"message_id,omitempty" xml:"message_id,omitempty"`
	// namespace
	Namespace string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}
