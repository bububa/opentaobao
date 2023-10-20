package product

import (
	"sync"
)

// GoodsPropertyOptionDto 结构体
type GoodsPropertyOptionDto struct {
	// 选项名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 属性值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 选项ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 属性ID
	AttrId int64 `json:"attr_id,omitempty" xml:"attr_id,omitempty"`
}

var poolGoodsPropertyOptionDto = sync.Pool{
	New: func() any {
		return new(GoodsPropertyOptionDto)
	},
}

// GetGoodsPropertyOptionDto() 从对象池中获取GoodsPropertyOptionDto
func GetGoodsPropertyOptionDto() *GoodsPropertyOptionDto {
	return poolGoodsPropertyOptionDto.Get().(*GoodsPropertyOptionDto)
}

// ReleaseGoodsPropertyOptionDto 释放GoodsPropertyOptionDto
func ReleaseGoodsPropertyOptionDto(v *GoodsPropertyOptionDto) {
	v.Name = ""
	v.Value = ""
	v.Id = 0
	v.AttrId = 0
	poolGoodsPropertyOptionDto.Put(v)
}
