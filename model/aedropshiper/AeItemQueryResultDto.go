package aedropshiper

import (
	"sync"
)

// AeItemQueryResultDto 结构体
type AeItemQueryResultDto struct {
	// SKU information
	AeItemSkuInfoDtos []AeItemSkuInfoDto `json:"ae_item_sku_info_dtos,omitempty" xml:"ae_item_sku_info_dtos>ae_item_sku_info_dto,omitempty"`
	// Attribute information
	AeItemProperties []AeItemProperty `json:"ae_item_properties,omitempty" xml:"ae_item_properties>ae_item_property,omitempty"`
	// Basic commodity information
	AeItemBaseInfoDto *AeItemBaseInfoDto `json:"ae_item_base_info_dto,omitempty" xml:"ae_item_base_info_dto,omitempty"`
	// Multimedia information
	AeMultimediaInfoDto *AeMultimediaInfoDto `json:"ae_multimedia_info_dto,omitempty" xml:"ae_multimedia_info_dto,omitempty"`
	// Package information
	PackageInfoDto *PackageInfoDto `json:"package_info_dto,omitempty" xml:"package_info_dto,omitempty"`
	// Logistics information
	LogisticsInfoDto *LogisticsInfoDto `json:"logistics_info_dto,omitempty" xml:"logistics_info_dto,omitempty"`
	// Store Information
	AeStoreInfo *AeStoreInfo `json:"ae_store_info,omitempty" xml:"ae_store_info,omitempty"`
	// product id converter result
	ProductIdConverterResult *ProductIdConverterResultDto `json:"product_id_converter_result,omitempty" xml:"product_id_converter_result,omitempty"`
}

var poolAeItemQueryResultDto = sync.Pool{
	New: func() any {
		return new(AeItemQueryResultDto)
	},
}

// GetAeItemQueryResultDto() 从对象池中获取AeItemQueryResultDto
func GetAeItemQueryResultDto() *AeItemQueryResultDto {
	return poolAeItemQueryResultDto.Get().(*AeItemQueryResultDto)
}

// ReleaseAeItemQueryResultDto 释放AeItemQueryResultDto
func ReleaseAeItemQueryResultDto(v *AeItemQueryResultDto) {
	v.AeItemSkuInfoDtos = v.AeItemSkuInfoDtos[:0]
	v.AeItemProperties = v.AeItemProperties[:0]
	v.AeItemBaseInfoDto = nil
	v.AeMultimediaInfoDto = nil
	v.PackageInfoDto = nil
	v.LogisticsInfoDto = nil
	v.AeStoreInfo = nil
	v.ProductIdConverterResult = nil
	poolAeItemQueryResultDto.Put(v)
}
