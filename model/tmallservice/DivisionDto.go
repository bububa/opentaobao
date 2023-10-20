package tmallservice

import (
	"sync"
)

// DivisionDto 结构体
type DivisionDto struct {
	// 1
	DivisionNames []string `json:"division_names,omitempty" xml:"division_names>string,omitempty"`
}

var poolDivisionDto = sync.Pool{
	New: func() any {
		return new(DivisionDto)
	},
}

// GetDivisionDto() 从对象池中获取DivisionDto
func GetDivisionDto() *DivisionDto {
	return poolDivisionDto.Get().(*DivisionDto)
}

// ReleaseDivisionDto 释放DivisionDto
func ReleaseDivisionDto(v *DivisionDto) {
	v.DivisionNames = v.DivisionNames[:0]
	poolDivisionDto.Put(v)
}
