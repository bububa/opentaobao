package ascp

import (
	"sync"
)

// BatchCreateItemMappingResponse 结构体
type BatchCreateItemMappingResponse struct {
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 业务处理结果
	Data *BatchCreateItemMappingResult `json:"data,omitempty" xml:"data,omitempty"`
	// 成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBatchCreateItemMappingResponse = sync.Pool{
	New: func() any {
		return new(BatchCreateItemMappingResponse)
	},
}

// GetBatchCreateItemMappingResponse() 从对象池中获取BatchCreateItemMappingResponse
func GetBatchCreateItemMappingResponse() *BatchCreateItemMappingResponse {
	return poolBatchCreateItemMappingResponse.Get().(*BatchCreateItemMappingResponse)
}

// ReleaseBatchCreateItemMappingResponse 释放BatchCreateItemMappingResponse
func ReleaseBatchCreateItemMappingResponse(v *BatchCreateItemMappingResponse) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolBatchCreateItemMappingResponse.Put(v)
}
