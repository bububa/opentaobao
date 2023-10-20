package ascp

import (
	"sync"
)

// DeleteScItemRequest 结构体
type DeleteScItemRequest struct {
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 业务请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolDeleteScItemRequest = sync.Pool{
	New: func() any {
		return new(DeleteScItemRequest)
	},
}

// GetDeleteScItemRequest() 从对象池中获取DeleteScItemRequest
func GetDeleteScItemRequest() *DeleteScItemRequest {
	return poolDeleteScItemRequest.Get().(*DeleteScItemRequest)
}

// ReleaseDeleteScItemRequest 释放DeleteScItemRequest
func ReleaseDeleteScItemRequest(v *DeleteScItemRequest) {
	v.ScItemCode = ""
	v.RequestId = ""
	v.RequestTime = 0
	poolDeleteScItemRequest.Put(v)
}
