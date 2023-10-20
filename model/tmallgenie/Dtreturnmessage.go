package tmallgenie

import (
	"sync"
)

// Dtreturnmessage 结构体
type Dtreturnmessage struct {
	// 标识
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolDtreturnmessage = sync.Pool{
	New: func() any {
		return new(Dtreturnmessage)
	},
}

// GetDtreturnmessage() 从对象池中获取Dtreturnmessage
func GetDtreturnmessage() *Dtreturnmessage {
	return poolDtreturnmessage.Get().(*Dtreturnmessage)
}

// ReleaseDtreturnmessage 释放Dtreturnmessage
func ReleaseDtreturnmessage(v *Dtreturnmessage) {
	v.Type = ""
	v.Message = ""
	poolDtreturnmessage.Put(v)
}
