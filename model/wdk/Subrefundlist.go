package wdk

import (
	"sync"
)

// Subrefundlist 结构体
type Subrefundlist struct {
	// 外部子订单ID
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 期望取货数量
	ExpectFetchQuantity int64 `json:"expect_fetch_quantity,omitempty" xml:"expect_fetch_quantity,omitempty"`
}

var poolSubrefundlist = sync.Pool{
	New: func() any {
		return new(Subrefundlist)
	},
}

// GetSubrefundlist() 从对象池中获取Subrefundlist
func GetSubrefundlist() *Subrefundlist {
	return poolSubrefundlist.Get().(*Subrefundlist)
}

// ReleaseSubrefundlist 释放Subrefundlist
func ReleaseSubrefundlist(v *Subrefundlist) {
	v.OutSubOrderId = ""
	v.RefundFee = 0
	v.ExpectFetchQuantity = 0
	poolSubrefundlist.Put(v)
}
