package cainiaohandover

import (
	"sync"
)

// OpenSenderParam 结构体
type OpenSenderParam struct {
	// 卖家后台地址id,用来获取卖家详细地址信息，传入值为Long型；
	SellerAddressId int64 `json:"seller_address_id,omitempty" xml:"seller_address_id,omitempty"`
}

var poolOpenSenderParam = sync.Pool{
	New: func() any {
		return new(OpenSenderParam)
	},
}

// GetOpenSenderParam() 从对象池中获取OpenSenderParam
func GetOpenSenderParam() *OpenSenderParam {
	return poolOpenSenderParam.Get().(*OpenSenderParam)
}

// ReleaseOpenSenderParam 释放OpenSenderParam
func ReleaseOpenSenderParam(v *OpenSenderParam) {
	v.SellerAddressId = 0
	poolOpenSenderParam.Put(v)
}
