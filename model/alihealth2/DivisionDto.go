package alihealth2

import (
	"sync"
)

// DivisionDto 结构体
type DivisionDto struct {
	// divisionName
	DivisionName string `json:"division_name,omitempty" xml:"division_name,omitempty"`
	// divisionId
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
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
	v.DivisionName = ""
	v.DivisionId = 0
	poolDivisionDto.Put(v)
}
