package wdklogistics

import (
	"sync"
)

// AlibabaWdkLogisticsPusPickupCararrivedData 结构体
type AlibabaWdkLogisticsPusPickupCararrivedData struct {
	// 自提点code
	StationCode string `json:"station_code,omitempty" xml:"station_code,omitempty"`
	// 脱敏订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolAlibabaWdkLogisticsPusPickupCararrivedData = sync.Pool{
	New: func() any {
		return new(AlibabaWdkLogisticsPusPickupCararrivedData)
	},
}

// GetAlibabaWdkLogisticsPusPickupCararrivedData() 从对象池中获取AlibabaWdkLogisticsPusPickupCararrivedData
func GetAlibabaWdkLogisticsPusPickupCararrivedData() *AlibabaWdkLogisticsPusPickupCararrivedData {
	return poolAlibabaWdkLogisticsPusPickupCararrivedData.Get().(*AlibabaWdkLogisticsPusPickupCararrivedData)
}

// ReleaseAlibabaWdkLogisticsPusPickupCararrivedData 释放AlibabaWdkLogisticsPusPickupCararrivedData
func ReleaseAlibabaWdkLogisticsPusPickupCararrivedData(v *AlibabaWdkLogisticsPusPickupCararrivedData) {
	v.StationCode = ""
	v.OrderCode = ""
	v.Status = ""
	poolAlibabaWdkLogisticsPusPickupCararrivedData.Put(v)
}
