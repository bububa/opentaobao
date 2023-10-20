package ieagency

import (
	"sync"
)

// RefundActivityVo 结构体
type RefundActivityVo struct {
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动收回金额(单位：分)
	TakeBackPrice int64 `json:"take_back_price,omitempty" xml:"take_back_price,omitempty"`
}

var poolRefundActivityVo = sync.Pool{
	New: func() any {
		return new(RefundActivityVo)
	},
}

// GetRefundActivityVo() 从对象池中获取RefundActivityVo
func GetRefundActivityVo() *RefundActivityVo {
	return poolRefundActivityVo.Get().(*RefundActivityVo)
}

// ReleaseRefundActivityVo 释放RefundActivityVo
func ReleaseRefundActivityVo(v *RefundActivityVo) {
	v.Name = ""
	v.TakeBackPrice = 0
	poolRefundActivityVo.Put(v)
}
