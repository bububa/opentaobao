package omniorder

import (
	"sync"
)

// OmniItemCategoryDto 结构体
type OmniItemCategoryDto struct {
	// props
	Props []OmniItemCategoryPropDto `json:"props,omitempty" xml:"props>omni_item_category_prop_dto,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目cid
	Cid int64 `json:"cid,omitempty" xml:"cid,omitempty"`
}

var poolOmniItemCategoryDto = sync.Pool{
	New: func() any {
		return new(OmniItemCategoryDto)
	},
}

// GetOmniItemCategoryDto() 从对象池中获取OmniItemCategoryDto
func GetOmniItemCategoryDto() *OmniItemCategoryDto {
	return poolOmniItemCategoryDto.Get().(*OmniItemCategoryDto)
}

// ReleaseOmniItemCategoryDto 释放OmniItemCategoryDto
func ReleaseOmniItemCategoryDto(v *OmniItemCategoryDto) {
	v.Props = v.Props[:0]
	v.CategoryName = ""
	v.Cid = 0
	poolOmniItemCategoryDto.Put(v)
}
