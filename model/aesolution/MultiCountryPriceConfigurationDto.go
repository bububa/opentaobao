package aesolution

import (
	"sync"
)

// MultiCountryPriceConfigurationDto 结构体
type MultiCountryPriceConfigurationDto struct {
	// Price list for different countries
	CountryPriceList []SingleCountryPriceDto `json:"country_price_list,omitempty" xml:"country_price_list>single_country_price_dto,omitempty"`
	// Currently supporting only absolute. Please test carefully before uploading products.
	PriceType string `json:"price_type,omitempty" xml:"price_type,omitempty"`
}

var poolMultiCountryPriceConfigurationDto = sync.Pool{
	New: func() any {
		return new(MultiCountryPriceConfigurationDto)
	},
}

// GetMultiCountryPriceConfigurationDto() 从对象池中获取MultiCountryPriceConfigurationDto
func GetMultiCountryPriceConfigurationDto() *MultiCountryPriceConfigurationDto {
	return poolMultiCountryPriceConfigurationDto.Get().(*MultiCountryPriceConfigurationDto)
}

// ReleaseMultiCountryPriceConfigurationDto 释放MultiCountryPriceConfigurationDto
func ReleaseMultiCountryPriceConfigurationDto(v *MultiCountryPriceConfigurationDto) {
	v.CountryPriceList = v.CountryPriceList[:0]
	v.PriceType = ""
	poolMultiCountryPriceConfigurationDto.Put(v)
}
