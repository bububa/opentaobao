package alicom

import (
	"sync"
)

// ProtocolList 结构体
type ProtocolList struct {
	// 协议
	Protocol string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

var poolProtocolList = sync.Pool{
	New: func() any {
		return new(ProtocolList)
	},
}

// GetProtocolList() 从对象池中获取ProtocolList
func GetProtocolList() *ProtocolList {
	return poolProtocolList.Get().(*ProtocolList)
}

// ReleaseProtocolList 释放ProtocolList
func ReleaseProtocolList(v *ProtocolList) {
	v.Protocol = ""
	poolProtocolList.Put(v)
}
