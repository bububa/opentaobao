package eticket

import (
	"sync"
)

// SendFailCallbackResp 结构体
type SendFailCallbackResp struct {
	// 回复业务KV
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
}

var poolSendFailCallbackResp = sync.Pool{
	New: func() any {
		return new(SendFailCallbackResp)
	},
}

// GetSendFailCallbackResp() 从对象池中获取SendFailCallbackResp
func GetSendFailCallbackResp() *SendFailCallbackResp {
	return poolSendFailCallbackResp.Get().(*SendFailCallbackResp)
}

// ReleaseSendFailCallbackResp 释放SendFailCallbackResp
func ReleaseSendFailCallbackResp(v *SendFailCallbackResp) {
	v.AttributeMap = ""
	poolSendFailCallbackResp.Put(v)
}
