package wdk

// SeriesSkuRequest 结构体
type SeriesSkuRequest struct {
	// 商品编码集合
	SkuCodes []string `json:"sku_codes,omitempty" xml:"sku_codes>string,omitempty"`
	// 默认商品编码
	DefaultSkuCode string `json:"default_sku_code,omitempty" xml:"default_sku_code,omitempty"`
	// 系列编码
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
	// 需要移除默认商品
	RemoveDefaultSku bool `json:"remove_default_sku,omitempty" xml:"remove_default_sku,omitempty"`
}
