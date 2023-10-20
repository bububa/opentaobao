package tmallsc

import (
	"sync"
)

// ServicePriceLine 结构体
type ServicePriceLine struct {
	// 服务外部编码
	ServiceOuterId string `json:"service_outer_id,omitempty" xml:"service_outer_id,omitempty"`
	// 原价
	ConsumerPrice int64 `json:"consumer_price,omitempty" xml:"consumer_price,omitempty"`
	// 教育优惠价格
	EduPrice int64 `json:"edu_price,omitempty" xml:"edu_price,omitempty"`
}

var poolServicePriceLine = sync.Pool{
	New: func() any {
		return new(ServicePriceLine)
	},
}

// GetServicePriceLine() 从对象池中获取ServicePriceLine
func GetServicePriceLine() *ServicePriceLine {
	return poolServicePriceLine.Get().(*ServicePriceLine)
}

// ReleaseServicePriceLine 释放ServicePriceLine
func ReleaseServicePriceLine(v *ServicePriceLine) {
	v.ServiceOuterId = ""
	v.ConsumerPrice = 0
	v.EduPrice = 0
	poolServicePriceLine.Put(v)
}
