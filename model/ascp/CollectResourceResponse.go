package ascp

import (
	"sync"
)

// CollectResourceResponse 结构体
type CollectResourceResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
}

var poolCollectResourceResponse = sync.Pool{
	New: func() any {
		return new(CollectResourceResponse)
	},
}

// GetCollectResourceResponse() 从对象池中获取CollectResourceResponse
func GetCollectResourceResponse() *CollectResourceResponse {
	return poolCollectResourceResponse.Get().(*CollectResourceResponse)
}

// ReleaseCollectResourceResponse 释放CollectResourceResponse
func ReleaseCollectResourceResponse(v *CollectResourceResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.IsRetry = false
	poolCollectResourceResponse.Put(v)
}
