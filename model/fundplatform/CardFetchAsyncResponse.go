package fundplatform

import (
	"sync"
)

// CardFetchAsyncResponse 结构体
type CardFetchAsyncResponse struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误消息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCardFetchAsyncResponse = sync.Pool{
	New: func() any {
		return new(CardFetchAsyncResponse)
	},
}

// GetCardFetchAsyncResponse() 从对象池中获取CardFetchAsyncResponse
func GetCardFetchAsyncResponse() *CardFetchAsyncResponse {
	return poolCardFetchAsyncResponse.Get().(*CardFetchAsyncResponse)
}

// ReleaseCardFetchAsyncResponse 释放CardFetchAsyncResponse
func ReleaseCardFetchAsyncResponse(v *CardFetchAsyncResponse) {
	v.ResultCode = ""
	v.ResultMessage = ""
	v.Success = false
	poolCardFetchAsyncResponse.Put(v)
}
