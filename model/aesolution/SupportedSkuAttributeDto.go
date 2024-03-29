package aesolution

import (
	"sync"
)

// SupportedSkuAttributeDto 结构体
type SupportedSkuAttributeDto struct {
	// aliexpress sku value list
	AliexpressSkuValueList []SkuValueSimplifiedInfoDto `json:"aliexpress_sku_value_list,omitempty" xml:"aliexpress_sku_value_list>sku_value_simplified_info_dto,omitempty"`
	// aliexpress sku name, the same field when indicating the sku_name for posting product
	AliexpressSkuName string `json:"aliexpress_sku_name,omitempty" xml:"aliexpress_sku_name,omitempty"`
	// Indicates whether this sku attribute is mandatory under this category
	Required bool `json:"required,omitempty" xml:"required,omitempty"`
	// whether the corresponding aliexpress_sku_name support customized name by merchants
	SupportCustomizedName bool `json:"support_customized_name,omitempty" xml:"support_customized_name,omitempty"`
	// whether the corresponding aliexpress_sku_name support customized picture
	SupportCustomizedPicture bool `json:"support_customized_picture,omitempty" xml:"support_customized_picture,omitempty"`
}

var poolSupportedSkuAttributeDto = sync.Pool{
	New: func() any {
		return new(SupportedSkuAttributeDto)
	},
}

// GetSupportedSkuAttributeDto() 从对象池中获取SupportedSkuAttributeDto
func GetSupportedSkuAttributeDto() *SupportedSkuAttributeDto {
	return poolSupportedSkuAttributeDto.Get().(*SupportedSkuAttributeDto)
}

// ReleaseSupportedSkuAttributeDto 释放SupportedSkuAttributeDto
func ReleaseSupportedSkuAttributeDto(v *SupportedSkuAttributeDto) {
	v.AliexpressSkuValueList = v.AliexpressSkuValueList[:0]
	v.AliexpressSkuName = ""
	v.Required = false
	v.SupportCustomizedName = false
	v.SupportCustomizedPicture = false
	poolSupportedSkuAttributeDto.Put(v)
}
