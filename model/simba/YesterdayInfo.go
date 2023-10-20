package simba

import (
	"sync"
)

// YesterdayInfo 结构体
type YesterdayInfo struct {
	// 昨日点击量
	Click string `json:"click,omitempty" xml:"click,omitempty"`
	// 昨日展现量
	Impression string `json:"impression,omitempty" xml:"impression,omitempty"`
}

var poolYesterdayInfo = sync.Pool{
	New: func() any {
		return new(YesterdayInfo)
	},
}

// GetYesterdayInfo() 从对象池中获取YesterdayInfo
func GetYesterdayInfo() *YesterdayInfo {
	return poolYesterdayInfo.Get().(*YesterdayInfo)
}

// ReleaseYesterdayInfo 释放YesterdayInfo
func ReleaseYesterdayInfo(v *YesterdayInfo) {
	v.Click = ""
	v.Impression = ""
	poolYesterdayInfo.Put(v)
}
