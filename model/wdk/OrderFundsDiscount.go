package wdk

import (
	"sync"
)

// OrderFundsDiscount 结构体
type OrderFundsDiscount struct {
	// 活动名称
	DiscountName string `json:"discount_name,omitempty" xml:"discount_name,omitempty"`
	// 活动类型
	DiscountType string `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 优惠金额(分)
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 商家优惠(分)
	DiscountMerchantFee int64 `json:"discount_merchant_fee,omitempty" xml:"discount_merchant_fee,omitempty"`
	// 平台优惠(分)
	DiscountPlatformFee int64 `json:"discount_platform_fee,omitempty" xml:"discount_platform_fee,omitempty"`
}

var poolOrderFundsDiscount = sync.Pool{
	New: func() any {
		return new(OrderFundsDiscount)
	},
}

// GetOrderFundsDiscount() 从对象池中获取OrderFundsDiscount
func GetOrderFundsDiscount() *OrderFundsDiscount {
	return poolOrderFundsDiscount.Get().(*OrderFundsDiscount)
}

// ReleaseOrderFundsDiscount 释放OrderFundsDiscount
func ReleaseOrderFundsDiscount(v *OrderFundsDiscount) {
	v.DiscountName = ""
	v.DiscountType = ""
	v.DiscountFee = 0
	v.DiscountMerchantFee = 0
	v.DiscountPlatformFee = 0
	poolOrderFundsDiscount.Put(v)
}
