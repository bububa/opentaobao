package charity

import (
	"sync"
)

// JumpAddressHsfResponse 结构体
type JumpAddressHsfResponse struct {
	// 跳转地址
	JumpUrl string `json:"jump_url,omitempty" xml:"jump_url,omitempty"`
}

var poolJumpAddressHsfResponse = sync.Pool{
	New: func() any {
		return new(JumpAddressHsfResponse)
	},
}

// GetJumpAddressHsfResponse() 从对象池中获取JumpAddressHsfResponse
func GetJumpAddressHsfResponse() *JumpAddressHsfResponse {
	return poolJumpAddressHsfResponse.Get().(*JumpAddressHsfResponse)
}

// ReleaseJumpAddressHsfResponse 释放JumpAddressHsfResponse
func ReleaseJumpAddressHsfResponse(v *JumpAddressHsfResponse) {
	v.JumpUrl = ""
	poolJumpAddressHsfResponse.Put(v)
}
