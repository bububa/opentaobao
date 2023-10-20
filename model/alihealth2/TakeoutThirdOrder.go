package alihealth2

import (
	"sync"
)

// TakeoutThirdOrder 结构体
type TakeoutThirdOrder struct {
	// 订单商品
	GoodsList []GoodsList `json:"goods_list,omitempty" xml:"goods_list>goods_list,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 订单失败原因
	FaildReason string `json:"faild_reason,omitempty" xml:"faild_reason,omitempty"`
	// 实际支付金额
	ActualPayFee int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
	// 店铺id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolTakeoutThirdOrder = sync.Pool{
	New: func() any {
		return new(TakeoutThirdOrder)
	},
}

// GetTakeoutThirdOrder() 从对象池中获取TakeoutThirdOrder
func GetTakeoutThirdOrder() *TakeoutThirdOrder {
	return poolTakeoutThirdOrder.Get().(*TakeoutThirdOrder)
}

// ReleaseTakeoutThirdOrder 释放TakeoutThirdOrder
func ReleaseTakeoutThirdOrder(v *TakeoutThirdOrder) {
	v.GoodsList = v.GoodsList[:0]
	v.PayTime = ""
	v.FaildReason = ""
	v.ActualPayFee = 0
	v.StoreId = 0
	poolTakeoutThirdOrder.Put(v)
}
