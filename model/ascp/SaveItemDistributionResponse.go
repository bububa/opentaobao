package ascp

import (
	"sync"
)

// SaveItemDistributionResponse 结构体
type SaveItemDistributionResponse struct {
	// 返回信息码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 操作码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 业务处理结果
	Data *DistributionResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 调用接口是否成功。调用成功之后，需要看data里面的success，才能知道业务处理是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolSaveItemDistributionResponse = sync.Pool{
	New: func() any {
		return new(SaveItemDistributionResponse)
	},
}

// GetSaveItemDistributionResponse() 从对象池中获取SaveItemDistributionResponse
func GetSaveItemDistributionResponse() *SaveItemDistributionResponse {
	return poolSaveItemDistributionResponse.Get().(*SaveItemDistributionResponse)
}

// ReleaseSaveItemDistributionResponse 释放SaveItemDistributionResponse
func ReleaseSaveItemDistributionResponse(v *SaveItemDistributionResponse) {
	v.Code = ""
	v.Message = ""
	v.TraceId = ""
	v.Data = nil
	v.Success = false
	poolSaveItemDistributionResponse.Put(v)
}
