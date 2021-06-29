package aesolution

// MultiCountryPriceConfigurationDto 
type MultiCountryPriceConfigurationDto struct {

    // Currently supporting only absolute. Please test carefully before uploading products.
    
    PriceType   string `json:"price_type,omitempty" xml:"price_type,omitempty"`
    

    // Price list for different countries
    
    CountryPriceList   []SingleCountryPriceDto `json:"country_price_list,omitempty" xml:"country_price_list,omitempty"`
    

}
