package cainiaohandover

import (
	"sync"
)

// OpenReturnerParam 结构体
type OpenReturnerParam struct {
	// 卖家后台地址id,用来获取卖家详细地址信息，传入值为Long型；
	SellerAddressId int64 `json:"seller_address_id,omitempty" xml:"seller_address_id,omitempty"`
}

var poolOpenReturnerParam = sync.Pool{
	New: func() any {
		return new(OpenReturnerParam)
	},
}

// GetOpenReturnerParam() 从对象池中获取OpenReturnerParam
func GetOpenReturnerParam() *OpenReturnerParam {
	return poolOpenReturnerParam.Get().(*OpenReturnerParam)
}

// ReleaseOpenReturnerParam 释放OpenReturnerParam
func ReleaseOpenReturnerParam(v *OpenReturnerParam) {
	v.SellerAddressId = 0
	poolOpenReturnerParam.Put(v)
}
