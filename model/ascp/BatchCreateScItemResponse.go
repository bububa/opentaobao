package ascp

import (
	"sync"
)

// BatchCreateScItemResponse 结构体
type BatchCreateScItemResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchCreateScItemResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchCreateScItemResponse = sync.Pool{
	New: func() any {
		return new(BatchCreateScItemResponse)
	},
}

// GetBatchCreateScItemResponse() 从对象池中获取BatchCreateScItemResponse
func GetBatchCreateScItemResponse() *BatchCreateScItemResponse {
	return poolBatchCreateScItemResponse.Get().(*BatchCreateScItemResponse)
}

// ReleaseBatchCreateScItemResponse 释放BatchCreateScItemResponse
func ReleaseBatchCreateScItemResponse(v *BatchCreateScItemResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolBatchCreateScItemResponse.Put(v)
}
