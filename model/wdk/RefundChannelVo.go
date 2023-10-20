package wdk

import (
	"sync"
)

// RefundChannelVo 结构体
type RefundChannelVo struct {
	// 渠道码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 渠道名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 付款吗
	PayCode string `json:"pay_code,omitempty" xml:"pay_code,omitempty"`
	// 渠道幂等ID
	UniqueId string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
	// 可退金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 已退金额 (单位分)
	RefundedAmount int64 `json:"refunded_amount,omitempty" xml:"refunded_amount,omitempty"`
	// 渠道退款状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolRefundChannelVo = sync.Pool{
	New: func() any {
		return new(RefundChannelVo)
	},
}

// GetRefundChannelVo() 从对象池中获取RefundChannelVo
func GetRefundChannelVo() *RefundChannelVo {
	return poolRefundChannelVo.Get().(*RefundChannelVo)
}

// ReleaseRefundChannelVo 释放RefundChannelVo
func ReleaseRefundChannelVo(v *RefundChannelVo) {
	v.Code = ""
	v.Name = ""
	v.PayCode = ""
	v.UniqueId = ""
	v.Amount = 0
	v.RefundedAmount = 0
	v.Status = 0
	poolRefundChannelVo.Put(v)
}
