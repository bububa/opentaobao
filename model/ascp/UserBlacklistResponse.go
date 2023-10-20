package ascp

import (
	"sync"
)

// UserBlacklistResponse 结构体
type UserBlacklistResponse struct {
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 系统成功失败   true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否可重试
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
}

var poolUserBlacklistResponse = sync.Pool{
	New: func() any {
		return new(UserBlacklistResponse)
	},
}

// GetUserBlacklistResponse() 从对象池中获取UserBlacklistResponse
func GetUserBlacklistResponse() *UserBlacklistResponse {
	return poolUserBlacklistResponse.Get().(*UserBlacklistResponse)
}

// ReleaseUserBlacklistResponse 释放UserBlacklistResponse
func ReleaseUserBlacklistResponse(v *UserBlacklistResponse) {
	v.Code = ""
	v.Message = ""
	v.Success = false
	v.IsRetry = false
	poolUserBlacklistResponse.Put(v)
}
