package ascp

import (
	"sync"
)

// BatchQueryConsignOrderResponse 结构体
type BatchQueryConsignOrderResponse struct {
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分页数据
	Data *PageData `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchQueryConsignOrderResponse = sync.Pool{
	New: func() any {
		return new(BatchQueryConsignOrderResponse)
	},
}

// GetBatchQueryConsignOrderResponse() 从对象池中获取BatchQueryConsignOrderResponse
func GetBatchQueryConsignOrderResponse() *BatchQueryConsignOrderResponse {
	return poolBatchQueryConsignOrderResponse.Get().(*BatchQueryConsignOrderResponse)
}

// ReleaseBatchQueryConsignOrderResponse 释放BatchQueryConsignOrderResponse
func ReleaseBatchQueryConsignOrderResponse(v *BatchQueryConsignOrderResponse) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolBatchQueryConsignOrderResponse.Put(v)
}
