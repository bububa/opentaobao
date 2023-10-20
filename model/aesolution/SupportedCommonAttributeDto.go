package aesolution

import (
	"sync"
)

// SupportedCommonAttributeDto 结构体
type SupportedCommonAttributeDto struct {
	// aliexpress common attribute value list
	AliexpressCommonAttributeValueList []CommonAttributeValueInfoDto `json:"aliexpress_common_attribute_value_list,omitempty" xml:"aliexpress_common_attribute_value_list>common_attribute_value_info_dto,omitempty"`
	// aliexpress common attribute name
	AliexpressCommonAttributeName string `json:"aliexpress_common_attribute_name,omitempty" xml:"aliexpress_common_attribute_name,omitempty"`
	// aliexpress common attribute name id
	AliexpressCommonAttributeNameId int64 `json:"aliexpress_common_attribute_name_id,omitempty" xml:"aliexpress_common_attribute_name_id,omitempty"`
	// whether the common attribute is required under this category
	Required bool `json:"required,omitempty" xml:"required,omitempty"`
}

var poolSupportedCommonAttributeDto = sync.Pool{
	New: func() any {
		return new(SupportedCommonAttributeDto)
	},
}

// GetSupportedCommonAttributeDto() 从对象池中获取SupportedCommonAttributeDto
func GetSupportedCommonAttributeDto() *SupportedCommonAttributeDto {
	return poolSupportedCommonAttributeDto.Get().(*SupportedCommonAttributeDto)
}

// ReleaseSupportedCommonAttributeDto 释放SupportedCommonAttributeDto
func ReleaseSupportedCommonAttributeDto(v *SupportedCommonAttributeDto) {
	v.AliexpressCommonAttributeValueList = v.AliexpressCommonAttributeValueList[:0]
	v.AliexpressCommonAttributeName = ""
	v.AliexpressCommonAttributeNameId = 0
	v.Required = false
	poolSupportedCommonAttributeDto.Put(v)
}
