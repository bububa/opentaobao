package aedropshiper

import (
	"sync"
)

// AeOrderLogisticsInfo 结构体
type AeOrderLogisticsInfo struct {
	// Logistics tracking number
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// Logistics Services
	LogisticsService string `json:"logistics_service,omitempty" xml:"logistics_service,omitempty"`
}

var poolAeOrderLogisticsInfo = sync.Pool{
	New: func() any {
		return new(AeOrderLogisticsInfo)
	},
}

// GetAeOrderLogisticsInfo() 从对象池中获取AeOrderLogisticsInfo
func GetAeOrderLogisticsInfo() *AeOrderLogisticsInfo {
	return poolAeOrderLogisticsInfo.Get().(*AeOrderLogisticsInfo)
}

// ReleaseAeOrderLogisticsInfo 释放AeOrderLogisticsInfo
func ReleaseAeOrderLogisticsInfo(v *AeOrderLogisticsInfo) {
	v.LogisticsNo = ""
	v.LogisticsService = ""
	poolAeOrderLogisticsInfo.Put(v)
}
