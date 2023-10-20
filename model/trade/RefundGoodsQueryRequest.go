package trade

import (
	"sync"
)

// RefundGoodsQueryRequest 结构体
type RefundGoodsQueryRequest struct {
	// 退货单id
	RefundGoodsId string `json:"refund_goods_id,omitempty" xml:"refund_goods_id,omitempty"`
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}

var poolRefundGoodsQueryRequest = sync.Pool{
	New: func() any {
		return new(RefundGoodsQueryRequest)
	},
}

// GetRefundGoodsQueryRequest() 从对象池中获取RefundGoodsQueryRequest
func GetRefundGoodsQueryRequest() *RefundGoodsQueryRequest {
	return poolRefundGoodsQueryRequest.Get().(*RefundGoodsQueryRequest)
}

// ReleaseRefundGoodsQueryRequest 释放RefundGoodsQueryRequest
func ReleaseRefundGoodsQueryRequest(v *RefundGoodsQueryRequest) {
	v.RefundGoodsId = ""
	v.ShopId = ""
	poolRefundGoodsQueryRequest.Put(v)
}
