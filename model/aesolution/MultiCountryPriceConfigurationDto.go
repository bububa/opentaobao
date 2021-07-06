package aesolution

// MultiCountryPriceConfigurationDto 结构体
type MultiCountryPriceConfigurationDto struct {
	// Price list for different countries
	CountryPriceList []SingleCountryPriceDto `json:"country_price_list,omitempty" xml:"country_price_list>single_country_price_dto,omitempty"`
	// Currently supporting only absolute. Please test carefully before uploading products.
	PriceType string `json:"price_type,omitempty" xml:"price_type,omitempty"`
}
