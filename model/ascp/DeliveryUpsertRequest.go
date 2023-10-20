package ascp

import (
	"sync"
)

// DeliveryUpsertRequest 结构体
type DeliveryUpsertRequest struct {
	// 快递信息数组；最多50条
	DeliveryInfos []DeliveryInfo `json:"delivery_infos,omitempty" xml:"delivery_infos>delivery_info,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolDeliveryUpsertRequest = sync.Pool{
	New: func() any {
		return new(DeliveryUpsertRequest)
	},
}

// GetDeliveryUpsertRequest() 从对象池中获取DeliveryUpsertRequest
func GetDeliveryUpsertRequest() *DeliveryUpsertRequest {
	return poolDeliveryUpsertRequest.Get().(*DeliveryUpsertRequest)
}

// ReleaseDeliveryUpsertRequest 释放DeliveryUpsertRequest
func ReleaseDeliveryUpsertRequest(v *DeliveryUpsertRequest) {
	v.DeliveryInfos = v.DeliveryInfos[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolDeliveryUpsertRequest.Put(v)
}
