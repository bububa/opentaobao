package baichuan

import (
	"sync"
)

// RichClientInfo 结构体
type RichClientInfo struct {
}

var poolRichClientInfo = sync.Pool{
	New: func() any {
		return new(RichClientInfo)
	},
}

// GetRichClientInfo() 从对象池中获取RichClientInfo
func GetRichClientInfo() *RichClientInfo {
	return poolRichClientInfo.Get().(*RichClientInfo)
}

// ReleaseRichClientInfo 释放RichClientInfo
func ReleaseRichClientInfo(v *RichClientInfo) {
	poolRichClientInfo.Put(v)
}
