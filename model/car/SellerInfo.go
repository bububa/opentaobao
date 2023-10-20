package car

import (
	"sync"
)

// SellerInfo 结构体
type SellerInfo struct {
	// sellerEmail
	SellerEmail string `json:"seller_email,omitempty" xml:"seller_email,omitempty"`
	// sellerNick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// sellerPhone
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// sellerId
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}

var poolSellerInfo = sync.Pool{
	New: func() any {
		return new(SellerInfo)
	},
}

// GetSellerInfo() 从对象池中获取SellerInfo
func GetSellerInfo() *SellerInfo {
	return poolSellerInfo.Get().(*SellerInfo)
}

// ReleaseSellerInfo 释放SellerInfo
func ReleaseSellerInfo(v *SellerInfo) {
	v.SellerEmail = ""
	v.SellerNick = ""
	v.SellerPhone = ""
	v.SellerId = 0
	poolSellerInfo.Put(v)
}
