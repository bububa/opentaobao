package ascp

import (
	"sync"
)

// DeliveryStatusUpdateRequest 结构体
type DeliveryStatusUpdateRequest struct {
	// 快递信息数组；最多50条
	DeliveryInfos []DeliveryInfo `json:"delivery_infos,omitempty" xml:"delivery_infos>delivery_info,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolDeliveryStatusUpdateRequest = sync.Pool{
	New: func() any {
		return new(DeliveryStatusUpdateRequest)
	},
}

// GetDeliveryStatusUpdateRequest() 从对象池中获取DeliveryStatusUpdateRequest
func GetDeliveryStatusUpdateRequest() *DeliveryStatusUpdateRequest {
	return poolDeliveryStatusUpdateRequest.Get().(*DeliveryStatusUpdateRequest)
}

// ReleaseDeliveryStatusUpdateRequest 释放DeliveryStatusUpdateRequest
func ReleaseDeliveryStatusUpdateRequest(v *DeliveryStatusUpdateRequest) {
	v.DeliveryInfos = v.DeliveryInfos[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolDeliveryStatusUpdateRequest.Put(v)
}
