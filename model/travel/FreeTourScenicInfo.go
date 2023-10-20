package travel

import (
	"sync"
)

// FreeTourScenicInfo 结构体
type FreeTourScenicInfo struct {
	// 必填，景点名称
	CnName string `json:"cn_name,omitempty" xml:"cn_name,omitempty"`
	// 必填，门票类型
	TicketType string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
	// 必填，景点所在城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
}

var poolFreeTourScenicInfo = sync.Pool{
	New: func() any {
		return new(FreeTourScenicInfo)
	},
}

// GetFreeTourScenicInfo() 从对象池中获取FreeTourScenicInfo
func GetFreeTourScenicInfo() *FreeTourScenicInfo {
	return poolFreeTourScenicInfo.Get().(*FreeTourScenicInfo)
}

// ReleaseFreeTourScenicInfo 释放FreeTourScenicInfo
func ReleaseFreeTourScenicInfo(v *FreeTourScenicInfo) {
	v.CnName = ""
	v.TicketType = ""
	v.City = ""
	poolFreeTourScenicInfo.Put(v)
}
