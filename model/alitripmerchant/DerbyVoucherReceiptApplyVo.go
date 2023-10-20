package alitripmerchant

import (
	"sync"
)

// DerbyVoucherReceiptApplyVo 结构体
type DerbyVoucherReceiptApplyVo struct {
	// 发票申请流水号
	FlowNumber string `json:"flow_number,omitempty" xml:"flow_number,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolDerbyVoucherReceiptApplyVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherReceiptApplyVo)
	},
}

// GetDerbyVoucherReceiptApplyVo() 从对象池中获取DerbyVoucherReceiptApplyVo
func GetDerbyVoucherReceiptApplyVo() *DerbyVoucherReceiptApplyVo {
	return poolDerbyVoucherReceiptApplyVo.Get().(*DerbyVoucherReceiptApplyVo)
}

// ReleaseDerbyVoucherReceiptApplyVo 释放DerbyVoucherReceiptApplyVo
func ReleaseDerbyVoucherReceiptApplyVo(v *DerbyVoucherReceiptApplyVo) {
	v.FlowNumber = ""
	v.OrderId = ""
	poolDerbyVoucherReceiptApplyVo.Put(v)
}
