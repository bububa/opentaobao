package charity

import (
	"sync"
)

// JumpAddressDto 结构体
type JumpAddressDto struct {
	// 跳转uri
	Uri string `json:"uri,omitempty" xml:"uri,omitempty"`
}

var poolJumpAddressDto = sync.Pool{
	New: func() any {
		return new(JumpAddressDto)
	},
}

// GetJumpAddressDto() 从对象池中获取JumpAddressDto
func GetJumpAddressDto() *JumpAddressDto {
	return poolJumpAddressDto.Get().(*JumpAddressDto)
}

// ReleaseJumpAddressDto 释放JumpAddressDto
func ReleaseJumpAddressDto(v *JumpAddressDto) {
	v.Uri = ""
	poolJumpAddressDto.Put(v)
}
