package trade

import (
	"sync"
)

// RefundGoodsSubOrder 结构体
type RefundGoodsSubOrder struct {
	// 退货商品子订单号
	SubBizOrderId string `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 退货件数
	GoodsAmount string `json:"goods_amount,omitempty" xml:"goods_amount,omitempty"`
	// 取货件数
	FulfillAmount string `json:"fulfill_amount,omitempty" xml:"fulfill_amount,omitempty"`
	// 商品skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 期待取货结束时间
	FulfillEndTime string `json:"fulfill_end_time,omitempty" xml:"fulfill_end_time,omitempty"`
	// 期待取货开始时间
	FulfillStartTime string `json:"fulfill_start_time,omitempty" xml:"fulfill_start_time,omitempty"`
	// 退款单id
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 退款金额
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 是否称重
	WeightItem bool `json:"weight_item,omitempty" xml:"weight_item,omitempty"`
	// 是否赠品
	Gift bool `json:"gift,omitempty" xml:"gift,omitempty"`
	// 是否离开货架
	LeftWarehouse bool `json:"left_warehouse,omitempty" xml:"left_warehouse,omitempty"`
}

var poolRefundGoodsSubOrder = sync.Pool{
	New: func() any {
		return new(RefundGoodsSubOrder)
	},
}

// GetRefundGoodsSubOrder() 从对象池中获取RefundGoodsSubOrder
func GetRefundGoodsSubOrder() *RefundGoodsSubOrder {
	return poolRefundGoodsSubOrder.Get().(*RefundGoodsSubOrder)
}

// ReleaseRefundGoodsSubOrder 释放RefundGoodsSubOrder
func ReleaseRefundGoodsSubOrder(v *RefundGoodsSubOrder) {
	v.SubBizOrderId = ""
	v.GoodsAmount = ""
	v.FulfillAmount = ""
	v.SkuCode = ""
	v.FulfillEndTime = ""
	v.FulfillStartTime = ""
	v.RefundId = ""
	v.RefundFee = 0
	v.WeightItem = false
	v.Gift = false
	v.LeftWarehouse = false
	poolRefundGoodsSubOrder.Put(v)
}
