package wdk

import (
	"sync"
)

// Skudetails 结构体
type Skudetails struct {
	// 履约子单id
	FulfillSubOrderId string `json:"fulfill_sub_order_id,omitempty" xml:"fulfill_sub_order_id,omitempty"`
	// 货品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
}

var poolSkudetails = sync.Pool{
	New: func() any {
		return new(Skudetails)
	},
}

// GetSkudetails() 从对象池中获取Skudetails
func GetSkudetails() *Skudetails {
	return poolSkudetails.Get().(*Skudetails)
}

// ReleaseSkudetails 释放Skudetails
func ReleaseSkudetails(v *Skudetails) {
	v.FulfillSubOrderId = ""
	v.SkuCode = ""
	poolSkudetails.Put(v)
}
