package aesolution

// MultiCountryPriceConfigurationDTO 
type MultiCountryPriceConfigurationDTO struct {
    // Currently supporting only absolute. Please test carefully before uploading products.
    PriceType   string `json:"price_type,omitempty" xml:"price_type,omitempty"`
    // Price list for different countries
    CountryPriceList   []SingleCountryPriceDTO `json:"country_price_list,omitempty" xml:"country_price_list>single_country_price_dto,omitempty"`
}
