package product

// TmallItemSkuSortGetApiResult 结构体
type TmallItemSkuSortGetApiResult struct {
	// 错误信息集合
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 属性排序结果信息
	Model []ItemSalePropSort `json:"model,omitempty" xml:"model>item_sale_prop_sort,omitempty"`
	// 执行结果
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}
