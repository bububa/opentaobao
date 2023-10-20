package tmallsc

import (
	"sync"
)

// IdentifyTaskDeliveryRequest 结构体
type IdentifyTaskDeliveryRequest struct {
	// 配送地纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 配送地经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 工单ID
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolIdentifyTaskDeliveryRequest = sync.Pool{
	New: func() any {
		return new(IdentifyTaskDeliveryRequest)
	},
}

// GetIdentifyTaskDeliveryRequest() 从对象池中获取IdentifyTaskDeliveryRequest
func GetIdentifyTaskDeliveryRequest() *IdentifyTaskDeliveryRequest {
	return poolIdentifyTaskDeliveryRequest.Get().(*IdentifyTaskDeliveryRequest)
}

// ReleaseIdentifyTaskDeliveryRequest 释放IdentifyTaskDeliveryRequest
func ReleaseIdentifyTaskDeliveryRequest(v *IdentifyTaskDeliveryRequest) {
	v.Latitude = ""
	v.Longitude = ""
	v.WorkcardId = 0
	poolIdentifyTaskDeliveryRequest.Put(v)
}
