package drugtrace

import (
	"sync"
)

// WesCodeRelationDto 结构体
type WesCodeRelationDto struct {
	// 存在上下级关系时返回下级码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 存在上下级关系时返回上级码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
}

var poolWesCodeRelationDto = sync.Pool{
	New: func() any {
		return new(WesCodeRelationDto)
	},
}

// GetWesCodeRelationDto() 从对象池中获取WesCodeRelationDto
func GetWesCodeRelationDto() *WesCodeRelationDto {
	return poolWesCodeRelationDto.Get().(*WesCodeRelationDto)
}

// ReleaseWesCodeRelationDto 释放WesCodeRelationDto
func ReleaseWesCodeRelationDto(v *WesCodeRelationDto) {
	v.Code = ""
	v.ParentCode = ""
	poolWesCodeRelationDto.Put(v)
}
