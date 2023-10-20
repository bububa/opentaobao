package idleisv

import (
	"sync"
)

// YoupinCpvResultDto 结构体
type YoupinCpvResultDto struct {
	// 属性值list
	ValueList []YoupinPropertyValueResultDto `json:"value_list,omitempty" xml:"value_list>youpin_property_value_result_dto,omitempty"`
	// 属性id
	PropertyId string `json:"property_id,omitempty" xml:"property_id,omitempty"`
	// 属性名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 是否包含子属性
	ShowSubProperty bool `json:"show_sub_property,omitempty" xml:"show_sub_property,omitempty"`
}

var poolYoupinCpvResultDto = sync.Pool{
	New: func() any {
		return new(YoupinCpvResultDto)
	},
}

// GetYoupinCpvResultDto() 从对象池中获取YoupinCpvResultDto
func GetYoupinCpvResultDto() *YoupinCpvResultDto {
	return poolYoupinCpvResultDto.Get().(*YoupinCpvResultDto)
}

// ReleaseYoupinCpvResultDto 释放YoupinCpvResultDto
func ReleaseYoupinCpvResultDto(v *YoupinCpvResultDto) {
	v.ValueList = v.ValueList[:0]
	v.PropertyId = ""
	v.Name = ""
	v.ShowSubProperty = false
	poolYoupinCpvResultDto.Put(v)
}
