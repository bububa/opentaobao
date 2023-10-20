package tmallgenie

import (
	"sync"
)

// Payload 结构体
type Payload struct {
	// 错误码，出错时返回
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息，出错时返回
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
}

var poolPayload = sync.Pool{
	New: func() any {
		return new(Payload)
	},
}

// GetPayload() 从对象池中获取Payload
func GetPayload() *Payload {
	return poolPayload.Get().(*Payload)
}

// ReleasePayload 释放Payload
func ReleasePayload(v *Payload) {
	v.ErrorCode = ""
	v.Message = ""
	v.DeviceId = ""
	poolPayload.Put(v)
}
