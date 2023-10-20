package fundplatform

import (
	"sync"
)

// CardActiveResponse 结构体
type CardActiveResponse struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误消息
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 是否处理成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCardActiveResponse = sync.Pool{
	New: func() any {
		return new(CardActiveResponse)
	},
}

// GetCardActiveResponse() 从对象池中获取CardActiveResponse
func GetCardActiveResponse() *CardActiveResponse {
	return poolCardActiveResponse.Get().(*CardActiveResponse)
}

// ReleaseCardActiveResponse 释放CardActiveResponse
func ReleaseCardActiveResponse(v *CardActiveResponse) {
	v.ResultCode = ""
	v.ResultMessage = ""
	v.Success = false
	poolCardActiveResponse.Put(v)
}
