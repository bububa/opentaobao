package aliexpresssumaitong

import (
	"sync"
)

// ExtraMap 结构体
type ExtraMap struct {
	// 支付收单号
	PaymentId string `json:"payment_id,omitempty" xml:"payment_id,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
}

var poolExtraMap = sync.Pool{
	New: func() any {
		return new(ExtraMap)
	},
}

// GetExtraMap() 从对象池中获取ExtraMap
func GetExtraMap() *ExtraMap {
	return poolExtraMap.Get().(*ExtraMap)
}

// ReleaseExtraMap 释放ExtraMap
func ReleaseExtraMap(v *ExtraMap) {
	v.PaymentId = ""
	v.OrderStatus = ""
	poolExtraMap.Put(v)
}
