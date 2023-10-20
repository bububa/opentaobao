package wdk

import (
	"sync"
)

// SubRefundConfirm 结构体
type SubRefundConfirm struct {
	// 退款Id
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 盒马子单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款邮费
	RefundPostFee int64 `json:"refund_post_fee,omitempty" xml:"refund_post_fee,omitempty"`
}

var poolSubRefundConfirm = sync.Pool{
	New: func() any {
		return new(SubRefundConfirm)
	},
}

// GetSubRefundConfirm() 从对象池中获取SubRefundConfirm
func GetSubRefundConfirm() *SubRefundConfirm {
	return poolSubRefundConfirm.Get().(*SubRefundConfirm)
}

// ReleaseSubRefundConfirm 释放SubRefundConfirm
func ReleaseSubRefundConfirm(v *SubRefundConfirm) {
	v.RefundId = ""
	v.SkuCode = ""
	v.SubBizOrderId = ""
	v.RefundFee = 0
	v.RefundPostFee = 0
	poolSubRefundConfirm.Put(v)
}
