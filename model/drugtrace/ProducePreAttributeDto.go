package drugtrace

import (
	"sync"
)

// ProducePreAttributeDto 结构体
type ProducePreAttributeDto struct {
	// 货品属性对象
	AttrInfoList []Attrinfolist `json:"attr_info_list,omitempty" xml:"attr_info_list>attrinfolist,omitempty"`
	// 属性规则-英文
	DefaultProducePreAttributeEn string `json:"default_produce_pre_attribute_en,omitempty" xml:"default_produce_pre_attribute_en,omitempty"`
	// 属性规则-中文
	DefaultProducePreAttribute string `json:"default_produce_pre_attribute,omitempty" xml:"default_produce_pre_attribute,omitempty"`
}

var poolProducePreAttributeDto = sync.Pool{
	New: func() any {
		return new(ProducePreAttributeDto)
	},
}

// GetProducePreAttributeDto() 从对象池中获取ProducePreAttributeDto
func GetProducePreAttributeDto() *ProducePreAttributeDto {
	return poolProducePreAttributeDto.Get().(*ProducePreAttributeDto)
}

// ReleaseProducePreAttributeDto 释放ProducePreAttributeDto
func ReleaseProducePreAttributeDto(v *ProducePreAttributeDto) {
	v.AttrInfoList = v.AttrInfoList[:0]
	v.DefaultProducePreAttributeEn = ""
	v.DefaultProducePreAttribute = ""
	poolProducePreAttributeDto.Put(v)
}
