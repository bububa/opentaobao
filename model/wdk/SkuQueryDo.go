package wdk

// SkuQueryDo 结构体
type SkuQueryDo struct {
	// 商品编码（多个条码用英文逗号隔开，最多支持20个）
	SkuCodes []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
	// 门店或DC编码,如果填写了渠道店id，该字段会被忽略
	OuCode string `json:"ou_code,omitempty" xml:"ou_code,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 渠道编码（默认4）,如果填写了渠道店id，该字段会被忽略
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
}
