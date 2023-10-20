package car

import (
	"sync"
)

// BuyerInfo 结构体
type BuyerInfo struct {
	// buyerEmail
	BuyerEmail string `json:"buyer_email,omitempty" xml:"buyer_email,omitempty"`
	// buyerPhone
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
}

var poolBuyerInfo = sync.Pool{
	New: func() any {
		return new(BuyerInfo)
	},
}

// GetBuyerInfo() 从对象池中获取BuyerInfo
func GetBuyerInfo() *BuyerInfo {
	return poolBuyerInfo.Get().(*BuyerInfo)
}

// ReleaseBuyerInfo 释放BuyerInfo
func ReleaseBuyerInfo(v *BuyerInfo) {
	v.BuyerEmail = ""
	v.BuyerPhone = ""
	poolBuyerInfo.Put(v)
}
