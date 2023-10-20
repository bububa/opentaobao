package pur

import (
	"sync"
)

// AccessProductAttrValueDto 结构体
type AccessProductAttrValueDto struct {
	// 属性值名称
	AttrValueNameList []string `json:"attr_value_name_list,omitempty" xml:"attr_value_name_list>string,omitempty"`
	// 属性Value值英文
	AttrEnValueNameList []string `json:"attr_en_value_name_list,omitempty" xml:"attr_en_value_name_list>string,omitempty"`
	// 属性名称
	AttrName string `json:"attr_name,omitempty" xml:"attr_name,omitempty"`
	// 属性Key英文值
	AttrEnName string `json:"attr_en_name,omitempty" xml:"attr_en_name,omitempty"`
	// 属性描述
	AttrDesc string `json:"attr_desc,omitempty" xml:"attr_desc,omitempty"`
	// 是否关键属性y/n
	IsKey string `json:"is_key,omitempty" xml:"is_key,omitempty"`
}

var poolAccessProductAttrValueDto = sync.Pool{
	New: func() any {
		return new(AccessProductAttrValueDto)
	},
}

// GetAccessProductAttrValueDto() 从对象池中获取AccessProductAttrValueDto
func GetAccessProductAttrValueDto() *AccessProductAttrValueDto {
	return poolAccessProductAttrValueDto.Get().(*AccessProductAttrValueDto)
}

// ReleaseAccessProductAttrValueDto 释放AccessProductAttrValueDto
func ReleaseAccessProductAttrValueDto(v *AccessProductAttrValueDto) {
	v.AttrValueNameList = v.AttrValueNameList[:0]
	v.AttrEnValueNameList = v.AttrEnValueNameList[:0]
	v.AttrName = ""
	v.AttrEnName = ""
	v.AttrDesc = ""
	v.IsKey = ""
	poolAccessProductAttrValueDto.Put(v)
}
