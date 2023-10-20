package ascp

import (
	"sync"
)

// BatchCreateCombineScItemResponse 结构体
type BatchCreateCombineScItemResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchCreateCombineScItemResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchCreateCombineScItemResponse = sync.Pool{
	New: func() any {
		return new(BatchCreateCombineScItemResponse)
	},
}

// GetBatchCreateCombineScItemResponse() 从对象池中获取BatchCreateCombineScItemResponse
func GetBatchCreateCombineScItemResponse() *BatchCreateCombineScItemResponse {
	return poolBatchCreateCombineScItemResponse.Get().(*BatchCreateCombineScItemResponse)
}

// ReleaseBatchCreateCombineScItemResponse 释放BatchCreateCombineScItemResponse
func ReleaseBatchCreateCombineScItemResponse(v *BatchCreateCombineScItemResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolBatchCreateCombineScItemResponse.Put(v)
}
