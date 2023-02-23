package aedropshiper

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
