package cainiaohandover

import (
	"sync"
)

// OpenPickupInfoParam 结构体
type OpenPickupInfoParam struct {
	// 卖家后台地址id,用来获取卖家详细地址信息，传入值为Long型；
	SellerAddressId int64 `json:"seller_address_id,omitempty" xml:"seller_address_id,omitempty"`
}

var poolOpenPickupInfoParam = sync.Pool{
	New: func() any {
		return new(OpenPickupInfoParam)
	},
}

// GetOpenPickupInfoParam() 从对象池中获取OpenPickupInfoParam
func GetOpenPickupInfoParam() *OpenPickupInfoParam {
	return poolOpenPickupInfoParam.Get().(*OpenPickupInfoParam)
}

// ReleaseOpenPickupInfoParam 释放OpenPickupInfoParam
func ReleaseOpenPickupInfoParam(v *OpenPickupInfoParam) {
	v.SellerAddressId = 0
	poolOpenPickupInfoParam.Put(v)
}
