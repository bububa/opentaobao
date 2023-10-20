package promotion

import (
	"sync"
)

// SellerGlobalDiscount 结构体
type SellerGlobalDiscount struct {
	// 折扣1折100，9折900
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolSellerGlobalDiscount = sync.Pool{
	New: func() any {
		return new(SellerGlobalDiscount)
	},
}

// GetSellerGlobalDiscount() 从对象池中获取SellerGlobalDiscount
func GetSellerGlobalDiscount() *SellerGlobalDiscount {
	return poolSellerGlobalDiscount.Get().(*SellerGlobalDiscount)
}

// ReleaseSellerGlobalDiscount 释放SellerGlobalDiscount
func ReleaseSellerGlobalDiscount(v *SellerGlobalDiscount) {
	v.Discount = 0
	poolSellerGlobalDiscount.Put(v)
}
