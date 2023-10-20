package ascp

import (
	"sync"
)

// QueryDeliveryTemplateRequest 结构体
type QueryDeliveryTemplateRequest struct {
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolQueryDeliveryTemplateRequest = sync.Pool{
	New: func() any {
		return new(QueryDeliveryTemplateRequest)
	},
}

// GetQueryDeliveryTemplateRequest() 从对象池中获取QueryDeliveryTemplateRequest
func GetQueryDeliveryTemplateRequest() *QueryDeliveryTemplateRequest {
	return poolQueryDeliveryTemplateRequest.Get().(*QueryDeliveryTemplateRequest)
}

// ReleaseQueryDeliveryTemplateRequest 释放QueryDeliveryTemplateRequest
func ReleaseQueryDeliveryTemplateRequest(v *QueryDeliveryTemplateRequest) {
	v.RequestId = ""
	v.RequestTime = 0
	poolQueryDeliveryTemplateRequest.Put(v)
}
