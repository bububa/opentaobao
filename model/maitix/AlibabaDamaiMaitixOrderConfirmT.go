package maitix

import (
	"sync"
)

// AlibabaDamaiMaitixOrderConfirmT 结构体
type AlibabaDamaiMaitixOrderConfirmT struct {
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 支付状态：0:失败,1:成功
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
}

var poolAlibabaDamaiMaitixOrderConfirmT = sync.Pool{
	New: func() any {
		return new(AlibabaDamaiMaitixOrderConfirmT)
	},
}

// GetAlibabaDamaiMaitixOrderConfirmT() 从对象池中获取AlibabaDamaiMaitixOrderConfirmT
func GetAlibabaDamaiMaitixOrderConfirmT() *AlibabaDamaiMaitixOrderConfirmT {
	return poolAlibabaDamaiMaitixOrderConfirmT.Get().(*AlibabaDamaiMaitixOrderConfirmT)
}

// ReleaseAlibabaDamaiMaitixOrderConfirmT 释放AlibabaDamaiMaitixOrderConfirmT
func ReleaseAlibabaDamaiMaitixOrderConfirmT(v *AlibabaDamaiMaitixOrderConfirmT) {
	v.OrderId = ""
	v.PayStatus = 0
	poolAlibabaDamaiMaitixOrderConfirmT.Put(v)
}
