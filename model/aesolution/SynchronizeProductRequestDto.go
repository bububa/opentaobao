package aesolution

import (
	"sync"
)

// SynchronizeProductRequestDto 结构体
type SynchronizeProductRequestDto struct {
	// The sku list, in which the inventory needs to be updated within the same product id. Maximum 200 skus.
	MultipleSkuUpdateList []SynchronizeSkuRequestDto `json:"multiple_sku_update_list,omitempty" xml:"multiple_sku_update_list>synchronize_sku_request_dto,omitempty"`
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// multi country price configuration
	MultiCountryPriceConfiguration *MultiCountryPriceConfigurationDto `json:"multi_country_price_configuration,omitempty" xml:"multi_country_price_configuration,omitempty"`
}

var poolSynchronizeProductRequestDto = sync.Pool{
	New: func() any {
		return new(SynchronizeProductRequestDto)
	},
}

// GetSynchronizeProductRequestDto() 从对象池中获取SynchronizeProductRequestDto
func GetSynchronizeProductRequestDto() *SynchronizeProductRequestDto {
	return poolSynchronizeProductRequestDto.Get().(*SynchronizeProductRequestDto)
}

// ReleaseSynchronizeProductRequestDto 释放SynchronizeProductRequestDto
func ReleaseSynchronizeProductRequestDto(v *SynchronizeProductRequestDto) {
	v.MultipleSkuUpdateList = v.MultipleSkuUpdateList[:0]
	v.ProductId = 0
	v.MultiCountryPriceConfiguration = nil
	poolSynchronizeProductRequestDto.Put(v)
}
