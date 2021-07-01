package wdk

// SkuSeriesEditResult 结构体
type SkuSeriesEditResult struct {
	// 成功的商品编码集
	FailedSkuCodes []string `json:"failed_sku_codes,omitempty" xml:"failed_sku_codes>string,omitempty"`
	// 失败的商品编码集
	SuccessedSkuCodes []string `json:"successed_sku_codes,omitempty" xml:"successed_sku_codes>string,omitempty"`
	// 系列编码
	SeriesId int64 `json:"series_id,omitempty" xml:"series_id,omitempty"`
}
