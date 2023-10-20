package icbudropshipping

import (
	"sync"
)

// DivisionInfoDto 结构体
type DivisionInfoDto struct {
	// City Code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// City Name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolDivisionInfoDto = sync.Pool{
	New: func() any {
		return new(DivisionInfoDto)
	},
}

// GetDivisionInfoDto() 从对象池中获取DivisionInfoDto
func GetDivisionInfoDto() *DivisionInfoDto {
	return poolDivisionInfoDto.Get().(*DivisionInfoDto)
}

// ReleaseDivisionInfoDto 释放DivisionInfoDto
func ReleaseDivisionInfoDto(v *DivisionInfoDto) {
	v.Code = ""
	v.Name = ""
	poolDivisionInfoDto.Put(v)
}
