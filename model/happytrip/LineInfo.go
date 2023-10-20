package happytrip

import (
	"sync"
)

// LineInfo 结构体
type LineInfo struct {
	// 出发位置名称
	FromLocationName string `json:"from_location_name,omitempty" xml:"from_location_name,omitempty"`
	// 目的位置名称
	ToLocationName string `json:"to_location_name,omitempty" xml:"to_location_name,omitempty"`
}

var poolLineInfo = sync.Pool{
	New: func() any {
		return new(LineInfo)
	},
}

// GetLineInfo() 从对象池中获取LineInfo
func GetLineInfo() *LineInfo {
	return poolLineInfo.Get().(*LineInfo)
}

// ReleaseLineInfo 释放LineInfo
func ReleaseLineInfo(v *LineInfo) {
	v.FromLocationName = ""
	v.ToLocationName = ""
	poolLineInfo.Put(v)
}
