package aliexpress

import (
	"sync"
)

// AliexpressFeedPostPublishResponse 结构体
type AliexpressFeedPostPublishResponse struct {
	// 成功返回结果，json字符串
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 是否请求成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliexpressFeedPostPublishResponse = sync.Pool{
	New: func() any {
		return new(AliexpressFeedPostPublishResponse)
	},
}

// GetAliexpressFeedPostPublishResponse() 从对象池中获取AliexpressFeedPostPublishResponse
func GetAliexpressFeedPostPublishResponse() *AliexpressFeedPostPublishResponse {
	return poolAliexpressFeedPostPublishResponse.Get().(*AliexpressFeedPostPublishResponse)
}

// ReleaseAliexpressFeedPostPublishResponse 释放AliexpressFeedPostPublishResponse
func ReleaseAliexpressFeedPostPublishResponse(v *AliexpressFeedPostPublishResponse) {
	v.Data = ""
	v.Message = ""
	v.Code = 0
	v.Success = false
	poolAliexpressFeedPostPublishResponse.Put(v)
}
