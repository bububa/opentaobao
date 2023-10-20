package idleisv

import (
	"sync"
)

// YoupinPropertyValueResultDto 结构体
type YoupinPropertyValueResultDto struct {
	// 属性值d
	ValueId string `json:"value_id,omitempty" xml:"value_id,omitempty"`
	// 属性值名称
	ValueName string `json:"value_name,omitempty" xml:"value_name,omitempty"`
}

var poolYoupinPropertyValueResultDto = sync.Pool{
	New: func() any {
		return new(YoupinPropertyValueResultDto)
	},
}

// GetYoupinPropertyValueResultDto() 从对象池中获取YoupinPropertyValueResultDto
func GetYoupinPropertyValueResultDto() *YoupinPropertyValueResultDto {
	return poolYoupinPropertyValueResultDto.Get().(*YoupinPropertyValueResultDto)
}

// ReleaseYoupinPropertyValueResultDto 释放YoupinPropertyValueResultDto
func ReleaseYoupinPropertyValueResultDto(v *YoupinPropertyValueResultDto) {
	v.ValueId = ""
	v.ValueName = ""
	poolYoupinPropertyValueResultDto.Put(v)
}
