package wdk

import (
	"sync"
)

// RefundApplyInfo 结构体
type RefundApplyInfo struct {
	// 逆向子单列表
	SubRefundOrders []SubRefundOrder `json:"sub_refund_orders,omitempty" xml:"sub_refund_orders>sub_refund_order,omitempty"`
	// 退款图片清单
	RefundPics []string `json:"refund_pics,omitempty" xml:"refund_pics>string,omitempty"`
	// 外部主单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 外部渠道店ID(与shop_id必选其一)
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部逆向单ID
	OutRefundId string `json:"out_refund_id,omitempty" xml:"out_refund_id,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 渠道店id(与out_shop_id必选其一)
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 退款备注、或问题描述等补充性文本
	RefundNote string `json:"refund_note,omitempty" xml:"refund_note,omitempty"`
	// 申请退款金额，单位：分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退的运费
	RefundPostFee int64 `json:"refund_post_fee,omitempty" xml:"refund_post_fee,omitempty"`
	// 退的包装费
	RefundPackageFee int64 `json:"refund_package_fee,omitempty" xml:"refund_package_fee,omitempty"`
	// 渠道来源(选填out_shop_id时该值必填)
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 退款类型，1:仅退款。2.仅退货。3.退货退款
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

var poolRefundApplyInfo = sync.Pool{
	New: func() any {
		return new(RefundApplyInfo)
	},
}

// GetRefundApplyInfo() 从对象池中获取RefundApplyInfo
func GetRefundApplyInfo() *RefundApplyInfo {
	return poolRefundApplyInfo.Get().(*RefundApplyInfo)
}

// ReleaseRefundApplyInfo 释放RefundApplyInfo
func ReleaseRefundApplyInfo(v *RefundApplyInfo) {
	v.SubRefundOrders = v.SubRefundOrders[:0]
	v.RefundPics = v.RefundPics[:0]
	v.OutOrderId = ""
	v.OutShopId = ""
	v.OutRefundId = ""
	v.RefundReason = ""
	v.ShopId = ""
	v.RefundNote = ""
	v.RefundFee = 0
	v.RefundPostFee = 0
	v.RefundPackageFee = 0
	v.OrderFrom = 0
	v.RefundType = 0
	poolRefundApplyInfo.Put(v)
}
