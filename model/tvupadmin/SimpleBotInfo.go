package tvupadmin

import (
	"sync"
)

// SimpleBotInfo 结构体
type SimpleBotInfo struct {
	// 设备名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 设备id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolSimpleBotInfo = sync.Pool{
	New: func() any {
		return new(SimpleBotInfo)
	},
}

// GetSimpleBotInfo() 从对象池中获取SimpleBotInfo
func GetSimpleBotInfo() *SimpleBotInfo {
	return poolSimpleBotInfo.Get().(*SimpleBotInfo)
}

// ReleaseSimpleBotInfo 释放SimpleBotInfo
func ReleaseSimpleBotInfo(v *SimpleBotInfo) {
	v.Name = ""
	v.Id = 0
	poolSimpleBotInfo.Put(v)
}
