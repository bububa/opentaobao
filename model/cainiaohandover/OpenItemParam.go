package cainiaohandover

// OpenItemParam 结构体
type OpenItemParam struct {
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品英文名称
	EnglishName string `json:"english_name,omitempty" xml:"english_name,omitempty"`
	// 商品本地名称
	LocalName string `json:"local_name,omitempty" xml:"local_name,omitempty"`
	// 商品单价，单位结算币种最小单位，如人民币分
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// sku
	Sku string `json:"sku,omitempty" xml:"sku,omitempty"`
	// 后台商品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 商品重量，单位g
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 商品属性，cf_normal：普货、cf_has_battery：含电。
	ItemFeatures []string `json:"item_features,omitempty" xml:"item_features>string,omitempty"`
	// 商品价格币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 商品总价
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 商品长度
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 商品宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 商品高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}
