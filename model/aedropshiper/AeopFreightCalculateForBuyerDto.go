package aedropshiper

// AeopFreightCalculateForBuyerDto 结构体
type AeopFreightCalculateForBuyerDto struct {
	// 城市编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 国家编码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 商品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 商品数量
	ProductNum int64 `json:"product_num,omitempty" xml:"product_num,omitempty"`
	// 省份编码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 发货国家
	SendGoodsCountryCode string `json:"send_goods_country_code,omitempty" xml:"send_goods_country_code,omitempty"`
	// 商品价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 商品价格币种
	PriceCurrency string `json:"price_currency,omitempty" xml:"price_currency,omitempty"`
}
