package simba

import (
	"sync"
)

// TopInfo 结构体
type TopInfo struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Ok bool `json:"ok,omitempty" xml:"ok,omitempty"`
}

var poolTopInfo = sync.Pool{
	New: func() any {
		return new(TopInfo)
	},
}

// GetTopInfo() 从对象池中获取TopInfo
func GetTopInfo() *TopInfo {
	return poolTopInfo.Get().(*TopInfo)
}

// ReleaseTopInfo 释放TopInfo
func ReleaseTopInfo(v *TopInfo) {
	v.Message = ""
	v.ErrorCode = ""
	v.Ok = false
	poolTopInfo.Put(v)
}
