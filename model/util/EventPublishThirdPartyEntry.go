package util

import (
	"sync"
)

// EventPublishThirdPartyEntry 结构体
type EventPublishThirdPartyEntry struct {
	// 事件编码
	TriggerCode string `json:"trigger_code,omitempty" xml:"trigger_code,omitempty"`
	// 事件内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

var poolEventPublishThirdPartyEntry = sync.Pool{
	New: func() any {
		return new(EventPublishThirdPartyEntry)
	},
}

// GetEventPublishThirdPartyEntry() 从对象池中获取EventPublishThirdPartyEntry
func GetEventPublishThirdPartyEntry() *EventPublishThirdPartyEntry {
	return poolEventPublishThirdPartyEntry.Get().(*EventPublishThirdPartyEntry)
}

// ReleaseEventPublishThirdPartyEntry 释放EventPublishThirdPartyEntry
func ReleaseEventPublishThirdPartyEntry(v *EventPublishThirdPartyEntry) {
	v.TriggerCode = ""
	v.Content = ""
	poolEventPublishThirdPartyEntry.Put(v)
}
