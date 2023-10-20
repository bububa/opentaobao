package product

// TmallItemSkuStatusGetApiResult 结构体
type TmallItemSkuStatusGetApiResult struct {
	// 错误信息
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 执行结果
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 返回内容
	ItemSkuStatus *ItemSkuStatus `json:"item_sku_status,omitempty" xml:"item_sku_status,omitempty"`
}
