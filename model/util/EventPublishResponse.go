package util

import (
	"sync"
)

// EventPublishResponse 结构体
type EventPublishResponse struct {
	// 流程返回值列表
	ResponseList []ProcessResponse `json:"response_list,omitempty" xml:"response_list>process_response,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 是否触发工作流
	ExecuteProcess bool `json:"execute_process,omitempty" xml:"execute_process,omitempty"`
}

var poolEventPublishResponse = sync.Pool{
	New: func() any {
		return new(EventPublishResponse)
	},
}

// GetEventPublishResponse() 从对象池中获取EventPublishResponse
func GetEventPublishResponse() *EventPublishResponse {
	return poolEventPublishResponse.Get().(*EventPublishResponse)
}

// ReleaseEventPublishResponse 释放EventPublishResponse
func ReleaseEventPublishResponse(v *EventPublishResponse) {
	v.ResponseList = v.ResponseList[:0]
	v.TraceId = ""
	v.ExecuteProcess = false
	poolEventPublishResponse.Put(v)
}
