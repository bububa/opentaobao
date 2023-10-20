package wdk

import (
	"sync"
)

// OrderDesensitizePhoneRequest 结构体
type OrderDesensitizePhoneRequest struct {
	// 业务单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 经营店编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}

var poolOrderDesensitizePhoneRequest = sync.Pool{
	New: func() any {
		return new(OrderDesensitizePhoneRequest)
	},
}

// GetOrderDesensitizePhoneRequest() 从对象池中获取OrderDesensitizePhoneRequest
func GetOrderDesensitizePhoneRequest() *OrderDesensitizePhoneRequest {
	return poolOrderDesensitizePhoneRequest.Get().(*OrderDesensitizePhoneRequest)
}

// ReleaseOrderDesensitizePhoneRequest 释放OrderDesensitizePhoneRequest
func ReleaseOrderDesensitizePhoneRequest(v *OrderDesensitizePhoneRequest) {
	v.BizOrderId = ""
	v.StoreCode = ""
	poolOrderDesensitizePhoneRequest.Put(v)
}
