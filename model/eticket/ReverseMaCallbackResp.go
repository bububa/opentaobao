package eticket

import (
	"sync"
)

// ReverseMaCallbackResp 结构体
type ReverseMaCallbackResp struct {
	// 业务参数KV
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
}

var poolReverseMaCallbackResp = sync.Pool{
	New: func() any {
		return new(ReverseMaCallbackResp)
	},
}

// GetReverseMaCallbackResp() 从对象池中获取ReverseMaCallbackResp
func GetReverseMaCallbackResp() *ReverseMaCallbackResp {
	return poolReverseMaCallbackResp.Get().(*ReverseMaCallbackResp)
}

// ReleaseReverseMaCallbackResp 释放ReverseMaCallbackResp
func ReleaseReverseMaCallbackResp(v *ReverseMaCallbackResp) {
	v.AttributeMap = ""
	poolReverseMaCallbackResp.Put(v)
}
