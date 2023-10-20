package aedropshiper

import (
	"sync"
)

// AeopOrderLogisticsInfo 结构体
type AeopOrderLogisticsInfo struct {
	// 物流追踪号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 物流服务
	LogisticsService string `json:"logistics_service,omitempty" xml:"logistics_service,omitempty"`
}

var poolAeopOrderLogisticsInfo = sync.Pool{
	New: func() any {
		return new(AeopOrderLogisticsInfo)
	},
}

// GetAeopOrderLogisticsInfo() 从对象池中获取AeopOrderLogisticsInfo
func GetAeopOrderLogisticsInfo() *AeopOrderLogisticsInfo {
	return poolAeopOrderLogisticsInfo.Get().(*AeopOrderLogisticsInfo)
}

// ReleaseAeopOrderLogisticsInfo 释放AeopOrderLogisticsInfo
func ReleaseAeopOrderLogisticsInfo(v *AeopOrderLogisticsInfo) {
	v.LogisticsNo = ""
	v.LogisticsService = ""
	poolAeopOrderLogisticsInfo.Put(v)
}
