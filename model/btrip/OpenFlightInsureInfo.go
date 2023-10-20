package btrip

import (
	"sync"
)

// OpenFlightInsureInfo 结构体
type OpenFlightInsureInfo struct {
	// 乘机人(保险人)姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 保单号
	InsureNo string `json:"insure_no,omitempty" xml:"insure_no,omitempty"`
	// 状态：1已出保 2已退保&#39;
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolOpenFlightInsureInfo = sync.Pool{
	New: func() any {
		return new(OpenFlightInsureInfo)
	},
}

// GetOpenFlightInsureInfo() 从对象池中获取OpenFlightInsureInfo
func GetOpenFlightInsureInfo() *OpenFlightInsureInfo {
	return poolOpenFlightInsureInfo.Get().(*OpenFlightInsureInfo)
}

// ReleaseOpenFlightInsureInfo 释放OpenFlightInsureInfo
func ReleaseOpenFlightInsureInfo(v *OpenFlightInsureInfo) {
	v.Name = ""
	v.InsureNo = ""
	v.Status = 0
	poolOpenFlightInsureInfo.Put(v)
}
