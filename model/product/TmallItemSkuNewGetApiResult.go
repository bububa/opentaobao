package product

// TmallItemSkuNewGetApiResult 结构体
type TmallItemSkuNewGetApiResult struct {
	// 错误码集合
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// sku标新信息集合
	Model []ItemSalePropNew `json:"model,omitempty" xml:"model>item_sale_prop_new,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
