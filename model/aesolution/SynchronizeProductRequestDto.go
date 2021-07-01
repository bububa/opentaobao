package aesolution

// SynchronizeProductRequestDto 结构体
type SynchronizeProductRequestDto struct {
	// product id
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// The sku list, in which the inventory needs to be updated within the same product id. Maximum 200 skus.
	MultipleSkuUpdateList []SynchronizeSkuRequestDto `json:"multiple_sku_update_list,omitempty" xml:"multiple_sku_update_list>synchronize_sku_request_dto,omitempty"`
	// multi country price configuration
	MultiCountryPriceConfiguration *MultiCountryPriceConfigurationDto `json:"multi_country_price_configuration,omitempty" xml:"multi_country_price_configuration,omitempty"`
}
