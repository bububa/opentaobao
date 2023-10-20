package omniorder

import (
	"sync"
)

// OmniItemCategoryPropDto 结构体
type OmniItemCategoryPropDto struct {
	// 销售属性名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// pid
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
}

var poolOmniItemCategoryPropDto = sync.Pool{
	New: func() any {
		return new(OmniItemCategoryPropDto)
	},
}

// GetOmniItemCategoryPropDto() 从对象池中获取OmniItemCategoryPropDto
func GetOmniItemCategoryPropDto() *OmniItemCategoryPropDto {
	return poolOmniItemCategoryPropDto.Get().(*OmniItemCategoryPropDto)
}

// ReleaseOmniItemCategoryPropDto 释放OmniItemCategoryPropDto
func ReleaseOmniItemCategoryPropDto(v *OmniItemCategoryPropDto) {
	v.Name = ""
	v.Pid = 0
	poolOmniItemCategoryPropDto.Put(v)
}
