package ascp

import (
	"sync"
)

// OrderReceiverPrivacyResponse 结构体
type OrderReceiverPrivacyResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 用户信息
	Data *OrderReceiverPrivacyData `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderReceiverPrivacyResponse = sync.Pool{
	New: func() any {
		return new(OrderReceiverPrivacyResponse)
	},
}

// GetOrderReceiverPrivacyResponse() 从对象池中获取OrderReceiverPrivacyResponse
func GetOrderReceiverPrivacyResponse() *OrderReceiverPrivacyResponse {
	return poolOrderReceiverPrivacyResponse.Get().(*OrderReceiverPrivacyResponse)
}

// ReleaseOrderReceiverPrivacyResponse 释放OrderReceiverPrivacyResponse
func ReleaseOrderReceiverPrivacyResponse(v *OrderReceiverPrivacyResponse) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolOrderReceiverPrivacyResponse.Put(v)
}
