package tmallgeniescp

import (
	"sync"
)

// AbstractRequest 结构体
type AbstractRequest struct {
	// 扩展参数
	RequestExtendJson string `json:"request_extend_json,omitempty" xml:"request_extend_json,omitempty"`
}

var poolAbstractRequest = sync.Pool{
	New: func() any {
		return new(AbstractRequest)
	},
}

// GetAbstractRequest() 从对象池中获取AbstractRequest
func GetAbstractRequest() *AbstractRequest {
	return poolAbstractRequest.Get().(*AbstractRequest)
}

// ReleaseAbstractRequest 释放AbstractRequest
func ReleaseAbstractRequest(v *AbstractRequest) {
	v.RequestExtendJson = ""
	poolAbstractRequest.Put(v)
}
