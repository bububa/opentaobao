package nlife

import (
	"sync"
)

// Refund 结构体
type Refund struct {
	// 退款渠道列表
	RefundBillList []FundBill `json:"refund_bill_list,omitempty" xml:"refund_bill_list>fund_bill,omitempty"`
	// refundTime
	RefundTime string `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	// outRequestNo
	OutRequestNo string `json:"out_request_no,omitempty" xml:"out_request_no,omitempty"`
	// 退货的商品，逗号分隔元素，商品和数量冒号分隔
	RefundGoods string `json:"refund_goods,omitempty" xml:"refund_goods,omitempty"`
}

var poolRefund = sync.Pool{
	New: func() any {
		return new(Refund)
	},
}

// GetRefund() 从对象池中获取Refund
func GetRefund() *Refund {
	return poolRefund.Get().(*Refund)
}

// ReleaseRefund 释放Refund
func ReleaseRefund(v *Refund) {
	v.RefundBillList = v.RefundBillList[:0]
	v.RefundTime = ""
	v.OutRequestNo = ""
	v.RefundGoods = ""
	poolRefund.Put(v)
}
