package simba

import (
	"sync"
)

// DimDtOs 结构体
type DimDtOs struct {
	// tagList
	TagList []TagOptions `json:"tag_list,omitempty" xml:"tag_list>tag_options,omitempty"`
	// 维度名称,如性别,年龄
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 维度id,如性别年龄的id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolDimDtOs = sync.Pool{
	New: func() any {
		return new(DimDtOs)
	},
}

// GetDimDtOs() 从对象池中获取DimDtOs
func GetDimDtOs() *DimDtOs {
	return poolDimDtOs.Get().(*DimDtOs)
}

// ReleaseDimDtOs 释放DimDtOs
func ReleaseDimDtOs(v *DimDtOs) {
	v.TagList = v.TagList[:0]
	v.Name = ""
	v.Id = 0
	poolDimDtOs.Put(v)
}
