package ascp

import (
	"sync"
)

// DeliveryStatusUpdateResponse 结构体
type DeliveryStatusUpdateResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 详细
	Data *DataDetail `json:"data,omitempty" xml:"data,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDeliveryStatusUpdateResponse = sync.Pool{
	New: func() any {
		return new(DeliveryStatusUpdateResponse)
	},
}

// GetDeliveryStatusUpdateResponse() 从对象池中获取DeliveryStatusUpdateResponse
func GetDeliveryStatusUpdateResponse() *DeliveryStatusUpdateResponse {
	return poolDeliveryStatusUpdateResponse.Get().(*DeliveryStatusUpdateResponse)
}

// ReleaseDeliveryStatusUpdateResponse 释放DeliveryStatusUpdateResponse
func ReleaseDeliveryStatusUpdateResponse(v *DeliveryStatusUpdateResponse) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolDeliveryStatusUpdateResponse.Put(v)
}
