package wdk

import (
	"sync"
)

// PaymentItemDo 结构体
type PaymentItemDo struct {
	// 商品sku码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品补差金额，单位为分
	PaymentFee int64 `json:"payment_fee,omitempty" xml:"payment_fee,omitempty"`
}

var poolPaymentItemDo = sync.Pool{
	New: func() any {
		return new(PaymentItemDo)
	},
}

// GetPaymentItemDo() 从对象池中获取PaymentItemDo
func GetPaymentItemDo() *PaymentItemDo {
	return poolPaymentItemDo.Get().(*PaymentItemDo)
}

// ReleasePaymentItemDo 释放PaymentItemDo
func ReleasePaymentItemDo(v *PaymentItemDo) {
	v.SkuCode = ""
	v.PaymentFee = 0
	poolPaymentItemDo.Put(v)
}
