package ascp

import (
	"sync"
)

// BatchUpdateScItemResponse 结构体
type BatchUpdateScItemResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchUpdateScItemResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchUpdateScItemResponse = sync.Pool{
	New: func() any {
		return new(BatchUpdateScItemResponse)
	},
}

// GetBatchUpdateScItemResponse() 从对象池中获取BatchUpdateScItemResponse
func GetBatchUpdateScItemResponse() *BatchUpdateScItemResponse {
	return poolBatchUpdateScItemResponse.Get().(*BatchUpdateScItemResponse)
}

// ReleaseBatchUpdateScItemResponse 释放BatchUpdateScItemResponse
func ReleaseBatchUpdateScItemResponse(v *BatchUpdateScItemResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolBatchUpdateScItemResponse.Put(v)
}
