package aesolution

import (
	"sync"
)

// CommonAttributeValueInfoDto 结构体
type CommonAttributeValueInfoDto struct {
	// aliexpress common attribute value
	AliexpressCommonAttributeValue string `json:"aliexpress_common_attribute_value,omitempty" xml:"aliexpress_common_attribute_value,omitempty"`
	// aliexpress common attribute value id
	AliexpressCommonAttributeValueId int64 `json:"aliexpress_common_attribute_value_id,omitempty" xml:"aliexpress_common_attribute_value_id,omitempty"`
}

var poolCommonAttributeValueInfoDto = sync.Pool{
	New: func() any {
		return new(CommonAttributeValueInfoDto)
	},
}

// GetCommonAttributeValueInfoDto() 从对象池中获取CommonAttributeValueInfoDto
func GetCommonAttributeValueInfoDto() *CommonAttributeValueInfoDto {
	return poolCommonAttributeValueInfoDto.Get().(*CommonAttributeValueInfoDto)
}

// ReleaseCommonAttributeValueInfoDto 释放CommonAttributeValueInfoDto
func ReleaseCommonAttributeValueInfoDto(v *CommonAttributeValueInfoDto) {
	v.AliexpressCommonAttributeValue = ""
	v.AliexpressCommonAttributeValueId = 0
	poolCommonAttributeValueInfoDto.Put(v)
}
