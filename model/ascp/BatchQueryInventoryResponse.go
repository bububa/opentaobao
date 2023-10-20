package ascp

import (
	"sync"
)

// BatchQueryInventoryResponse 结构体
type BatchQueryInventoryResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchQueryInventoryResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchQueryInventoryResponse = sync.Pool{
	New: func() any {
		return new(BatchQueryInventoryResponse)
	},
}

// GetBatchQueryInventoryResponse() 从对象池中获取BatchQueryInventoryResponse
func GetBatchQueryInventoryResponse() *BatchQueryInventoryResponse {
	return poolBatchQueryInventoryResponse.Get().(*BatchQueryInventoryResponse)
}

// ReleaseBatchQueryInventoryResponse 释放BatchQueryInventoryResponse
func ReleaseBatchQueryInventoryResponse(v *BatchQueryInventoryResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolBatchQueryInventoryResponse.Put(v)
}
