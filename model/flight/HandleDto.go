package flight

import (
	"sync"
)

// HandleDto 结构体
type HandleDto struct {
	// 跟进人名称
	CurrentFollowerName string `json:"current_follower_name,omitempty" xml:"current_follower_name,omitempty"`
	// 协同单id
	CaseId int64 `json:"case_id,omitempty" xml:"case_id,omitempty"`
	// 1:国内，2:国际
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}

var poolHandleDto = sync.Pool{
	New: func() any {
		return new(HandleDto)
	},
}

// GetHandleDto() 从对象池中获取HandleDto
func GetHandleDto() *HandleDto {
	return poolHandleDto.Get().(*HandleDto)
}

// ReleaseHandleDto 释放HandleDto
func ReleaseHandleDto(v *HandleDto) {
	v.CurrentFollowerName = ""
	v.CaseId = 0
	v.DomesticIntl = 0
	poolHandleDto.Put(v)
}
