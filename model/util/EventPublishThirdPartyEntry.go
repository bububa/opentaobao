package util

// EventPublishThirdPartyEntry 结构体
type EventPublishThirdPartyEntry struct {
	// 事件编码
	TriggerCode string `json:"trigger_code,omitempty" xml:"trigger_code,omitempty"`
	// 事件内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}
