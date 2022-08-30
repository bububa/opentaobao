package pur

// AccessQuotationDto 结构体
type AccessQuotationDto struct {
	// 阶梯价信息
	LadderPriceList []AccessLadderPriceDto `json:"ladder_price_list,omitempty" xml:"ladder_price_list>access_ladder_price_dto,omitempty"`
	// sku属性列表
	SkuAttrValueList []AccessSkuAttrValueDto `json:"sku_attr_value_list,omitempty" xml:"sku_attr_value_list>access_sku_attr_value_dto,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 报价生效时间
	EffectiveDate string `json:"effective_date,omitempty" xml:"effective_date,omitempty"`
	// 报价失效时间
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 外部商家skuId
	SourceSkuId string `json:"source_sku_id,omitempty" xml:"source_sku_id,omitempty"`
	// 协议价
	UnitPrice *BigDecimal `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 原价
	OriginUnitPrice *BigDecimal `json:"origin_unit_price,omitempty" xml:"origin_unit_price,omitempty"`
	// 是否阶梯价
	LadderPrice bool `json:"ladder_price,omitempty" xml:"ladder_price,omitempty"`
}
