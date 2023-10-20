package ascp

import (
	"sync"
)

// DeliveryResourceCreateResponse 结构体
type DeliveryResourceCreateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeliveryResourceCreateResponse = sync.Pool{
	New: func() any {
		return new(DeliveryResourceCreateResponse)
	},
}

// GetDeliveryResourceCreateResponse() 从对象池中获取DeliveryResourceCreateResponse
func GetDeliveryResourceCreateResponse() *DeliveryResourceCreateResponse {
	return poolDeliveryResourceCreateResponse.Get().(*DeliveryResourceCreateResponse)
}

// ReleaseDeliveryResourceCreateResponse 释放DeliveryResourceCreateResponse
func ReleaseDeliveryResourceCreateResponse(v *DeliveryResourceCreateResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	poolDeliveryResourceCreateResponse.Put(v)
}
