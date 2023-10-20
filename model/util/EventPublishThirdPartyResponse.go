package util

import (
	"sync"
)

// EventPublishThirdPartyResponse 结构体
type EventPublishThirdPartyResponse struct {
	// 发布事件结果列表
	EntryList []EventPublishThirdPartyResultEntry `json:"entry_list,omitempty" xml:"entry_list>event_publish_third_party_result_entry,omitempty"`
}

var poolEventPublishThirdPartyResponse = sync.Pool{
	New: func() any {
		return new(EventPublishThirdPartyResponse)
	},
}

// GetEventPublishThirdPartyResponse() 从对象池中获取EventPublishThirdPartyResponse
func GetEventPublishThirdPartyResponse() *EventPublishThirdPartyResponse {
	return poolEventPublishThirdPartyResponse.Get().(*EventPublishThirdPartyResponse)
}

// ReleaseEventPublishThirdPartyResponse 释放EventPublishThirdPartyResponse
func ReleaseEventPublishThirdPartyResponse(v *EventPublishThirdPartyResponse) {
	v.EntryList = v.EntryList[:0]
	poolEventPublishThirdPartyResponse.Put(v)
}
