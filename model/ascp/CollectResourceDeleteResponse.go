package ascp

import (
	"sync"
)

// CollectResourceDeleteResponse 结构体
type CollectResourceDeleteResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
}

var poolCollectResourceDeleteResponse = sync.Pool{
	New: func() any {
		return new(CollectResourceDeleteResponse)
	},
}

// GetCollectResourceDeleteResponse() 从对象池中获取CollectResourceDeleteResponse
func GetCollectResourceDeleteResponse() *CollectResourceDeleteResponse {
	return poolCollectResourceDeleteResponse.Get().(*CollectResourceDeleteResponse)
}

// ReleaseCollectResourceDeleteResponse 释放CollectResourceDeleteResponse
func ReleaseCollectResourceDeleteResponse(v *CollectResourceDeleteResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.IsRetry = false
	poolCollectResourceDeleteResponse.Put(v)
}
