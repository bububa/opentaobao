package alicom

import (
	"sync"
)

// CreatePhoneOrderReq 结构体
type CreatePhoneOrderReq struct {
	// 业务办理账户
	Account string `json:"account,omitempty" xml:"account,omitempty"`
	// 商家幂等ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolCreatePhoneOrderReq = sync.Pool{
	New: func() any {
		return new(CreatePhoneOrderReq)
	},
}

// GetCreatePhoneOrderReq() 从对象池中获取CreatePhoneOrderReq
func GetCreatePhoneOrderReq() *CreatePhoneOrderReq {
	return poolCreatePhoneOrderReq.Get().(*CreatePhoneOrderReq)
}

// ReleaseCreatePhoneOrderReq 释放CreatePhoneOrderReq
func ReleaseCreatePhoneOrderReq(v *CreatePhoneOrderReq) {
	v.Account = ""
	v.OutOrderId = ""
	v.ItemId = 0
	poolCreatePhoneOrderReq.Put(v)
}
