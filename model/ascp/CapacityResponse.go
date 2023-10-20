package ascp

import (
	"sync"
)

// CapacityResponse 结构体
type CapacityResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
}

var poolCapacityResponse = sync.Pool{
	New: func() any {
		return new(CapacityResponse)
	},
}

// GetCapacityResponse() 从对象池中获取CapacityResponse
func GetCapacityResponse() *CapacityResponse {
	return poolCapacityResponse.Get().(*CapacityResponse)
}

// ReleaseCapacityResponse 释放CapacityResponse
func ReleaseCapacityResponse(v *CapacityResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.IsRetry = false
	poolCapacityResponse.Put(v)
}
