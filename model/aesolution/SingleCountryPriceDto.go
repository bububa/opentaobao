package aesolution

import (
	"sync"
)

// SingleCountryPriceDto 结构体
type SingleCountryPriceDto struct {
	// Sku price list under the same ship_to_country
	SkuPriceByCountryList []SingleSkuPriceByCountryDto `json:"sku_price_by_country_list,omitempty" xml:"sku_price_by_country_list>single_sku_price_by_country_dto,omitempty"`
	// Currently supporting RU US CA ES FR UK NL IL BR CL AU UA BY JP TH SG KR ID MY PH VN IT DE SA AE PL TR
	ShipToCountry string `json:"ship_to_country,omitempty" xml:"ship_to_country,omitempty"`
}

var poolSingleCountryPriceDto = sync.Pool{
	New: func() any {
		return new(SingleCountryPriceDto)
	},
}

// GetSingleCountryPriceDto() 从对象池中获取SingleCountryPriceDto
func GetSingleCountryPriceDto() *SingleCountryPriceDto {
	return poolSingleCountryPriceDto.Get().(*SingleCountryPriceDto)
}

// ReleaseSingleCountryPriceDto 释放SingleCountryPriceDto
func ReleaseSingleCountryPriceDto(v *SingleCountryPriceDto) {
	v.SkuPriceByCountryList = v.SkuPriceByCountryList[:0]
	v.ShipToCountry = ""
	poolSingleCountryPriceDto.Put(v)
}
