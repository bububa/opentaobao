package ascp

import (
	"sync"
)

// TopResponse 结构体
type TopResponse struct {
	// 操作码
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求结果
	Data *QueryDistributorResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopResponse = sync.Pool{
	New: func() any {
		return new(TopResponse)
	},
}

// GetTopResponse() 从对象池中获取TopResponse
func GetTopResponse() *TopResponse {
	return poolTopResponse.Get().(*TopResponse)
}

// ReleaseTopResponse 释放TopResponse
func ReleaseTopResponse(v *TopResponse) {
	v.TraceId = ""
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.Success = false
	poolTopResponse.Put(v)
}
