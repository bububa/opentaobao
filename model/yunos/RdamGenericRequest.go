package yunos

import (
	"sync"
)

// RdamGenericRequest 结构体
type RdamGenericRequest struct {
	// 请求标示
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
}

var poolRdamGenericRequest = sync.Pool{
	New: func() any {
		return new(RdamGenericRequest)
	},
}

// GetRdamGenericRequest() 从对象池中获取RdamGenericRequest
func GetRdamGenericRequest() *RdamGenericRequest {
	return poolRdamGenericRequest.Get().(*RdamGenericRequest)
}

// ReleaseRdamGenericRequest 释放RdamGenericRequest
func ReleaseRdamGenericRequest(v *RdamGenericRequest) {
	v.TraceId = ""
	poolRdamGenericRequest.Put(v)
}
