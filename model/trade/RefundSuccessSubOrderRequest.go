package trade

import (
	"sync"
)

// RefundSuccessSubOrderRequest 结构体
type RefundSuccessSubOrderRequest struct {
	// 子订单id
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 退货单id
	RefundGoodsId string `json:"refund_goods_id,omitempty" xml:"refund_goods_id,omitempty"`
	// 履约单id
	FulfillId string `json:"fulfill_id,omitempty" xml:"fulfill_id,omitempty"`
	// 退款单id
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 实际退货数量
	ActualRefundQuantity string `json:"actual_refund_quantity,omitempty" xml:"actual_refund_quantity,omitempty"`
	// 实际退款金额
	RefundAmount string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 期望取货销售数量
	ExpectFetchSaleQuantity string `json:"expect_fetch_sale_quantity,omitempty" xml:"expect_fetch_sale_quantity,omitempty"`
	// 退货状态(1:退货退款；0:仅退款)
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

var poolRefundSuccessSubOrderRequest = sync.Pool{
	New: func() any {
		return new(RefundSuccessSubOrderRequest)
	},
}

// GetRefundSuccessSubOrderRequest() 从对象池中获取RefundSuccessSubOrderRequest
func GetRefundSuccessSubOrderRequest() *RefundSuccessSubOrderRequest {
	return poolRefundSuccessSubOrderRequest.Get().(*RefundSuccessSubOrderRequest)
}

// ReleaseRefundSuccessSubOrderRequest 释放RefundSuccessSubOrderRequest
func ReleaseRefundSuccessSubOrderRequest(v *RefundSuccessSubOrderRequest) {
	v.SubBizOrderId = ""
	v.RefundGoodsId = ""
	v.FulfillId = ""
	v.RefundId = ""
	v.SkuCode = ""
	v.ActualRefundQuantity = ""
	v.RefundAmount = ""
	v.ExpectFetchSaleQuantity = ""
	v.RefundType = 0
	poolRefundSuccessSubOrderRequest.Put(v)
}
