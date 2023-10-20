package fundplatform

import (
	"sync"
)

// CardMakingInformResponse 结构体
type CardMakingInformResponse struct {
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCardMakingInformResponse = sync.Pool{
	New: func() any {
		return new(CardMakingInformResponse)
	},
}

// GetCardMakingInformResponse() 从对象池中获取CardMakingInformResponse
func GetCardMakingInformResponse() *CardMakingInformResponse {
	return poolCardMakingInformResponse.Get().(*CardMakingInformResponse)
}

// ReleaseCardMakingInformResponse 释放CardMakingInformResponse
func ReleaseCardMakingInformResponse(v *CardMakingInformResponse) {
	v.Success = false
	poolCardMakingInformResponse.Put(v)
}
