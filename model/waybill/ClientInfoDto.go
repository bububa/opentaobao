package waybill

import (
	"sync"
)

// ClientInfoDto 结构体
type ClientInfoDto struct {
	// 调用时自定义描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}

var poolClientInfoDto = sync.Pool{
	New: func() any {
		return new(ClientInfoDto)
	},
}

// GetClientInfoDto() 从对象池中获取ClientInfoDto
func GetClientInfoDto() *ClientInfoDto {
	return poolClientInfoDto.Get().(*ClientInfoDto)
}

// ReleaseClientInfoDto 释放ClientInfoDto
func ReleaseClientInfoDto(v *ClientInfoDto) {
	v.Description = ""
	poolClientInfoDto.Put(v)
}
