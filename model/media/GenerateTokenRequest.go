package media

import (
	"sync"
)

// GenerateTokenRequest 结构体
type GenerateTokenRequest struct {
	// 请求策略
	UploadPolicy *UploadPolicy `json:"upload_policy,omitempty" xml:"upload_policy,omitempty"`
}

var poolGenerateTokenRequest = sync.Pool{
	New: func() any {
		return new(GenerateTokenRequest)
	},
}

// GetGenerateTokenRequest() 从对象池中获取GenerateTokenRequest
func GetGenerateTokenRequest() *GenerateTokenRequest {
	return poolGenerateTokenRequest.Get().(*GenerateTokenRequest)
}

// ReleaseGenerateTokenRequest 释放GenerateTokenRequest
func ReleaseGenerateTokenRequest(v *GenerateTokenRequest) {
	v.UploadPolicy = nil
	poolGenerateTokenRequest.Put(v)
}
