package util

import (
	"sync"
)

// EventPublishThirdPartyResultEntry 结构体
type EventPublishThirdPartyResultEntry struct {
	// 事件编码
	TriggerCode string `json:"trigger_code,omitempty" xml:"trigger_code,omitempty"`
	// 事件唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	SubErrCode string `json:"sub_err_code,omitempty" xml:"sub_err_code,omitempty"`
	// 错误内容
	SubErrMsg string `json:"sub_err_msg,omitempty" xml:"sub_err_msg,omitempty"`
	// 发布是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolEventPublishThirdPartyResultEntry = sync.Pool{
	New: func() any {
		return new(EventPublishThirdPartyResultEntry)
	},
}

// GetEventPublishThirdPartyResultEntry() 从对象池中获取EventPublishThirdPartyResultEntry
func GetEventPublishThirdPartyResultEntry() *EventPublishThirdPartyResultEntry {
	return poolEventPublishThirdPartyResultEntry.Get().(*EventPublishThirdPartyResultEntry)
}

// ReleaseEventPublishThirdPartyResultEntry 释放EventPublishThirdPartyResultEntry
func ReleaseEventPublishThirdPartyResultEntry(v *EventPublishThirdPartyResultEntry) {
	v.TriggerCode = ""
	v.TraceId = ""
	v.SubErrCode = ""
	v.SubErrMsg = ""
	v.Success = false
	poolEventPublishThirdPartyResultEntry.Put(v)
}
