package ascp

import (
	"sync"
)

// BatchUpdateCombineScItemResponse 结构体
type BatchUpdateCombineScItemResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchUpdateCombineScItemResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchUpdateCombineScItemResponse = sync.Pool{
	New: func() any {
		return new(BatchUpdateCombineScItemResponse)
	},
}

// GetBatchUpdateCombineScItemResponse() 从对象池中获取BatchUpdateCombineScItemResponse
func GetBatchUpdateCombineScItemResponse() *BatchUpdateCombineScItemResponse {
	return poolBatchUpdateCombineScItemResponse.Get().(*BatchUpdateCombineScItemResponse)
}

// ReleaseBatchUpdateCombineScItemResponse 释放BatchUpdateCombineScItemResponse
func ReleaseBatchUpdateCombineScItemResponse(v *BatchUpdateCombineScItemResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolBatchUpdateCombineScItemResponse.Put(v)
}
