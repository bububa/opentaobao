package omniorder

import (
	"sync"
)

// TpDto 结构体
type TpDto struct {
	// 公司编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 公司名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolTpDto = sync.Pool{
	New: func() any {
		return new(TpDto)
	},
}

// GetTpDto() 从对象池中获取TpDto
func GetTpDto() *TpDto {
	return poolTpDto.Get().(*TpDto)
}

// ReleaseTpDto 释放TpDto
func ReleaseTpDto(v *TpDto) {
	v.Code = ""
	v.Name = ""
	poolTpDto.Put(v)
}
