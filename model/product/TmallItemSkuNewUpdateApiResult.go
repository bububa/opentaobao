package product

// TmallItemSkuNewUpdateApiResult 结构体
type TmallItemSkuNewUpdateApiResult struct {
	// 错误码信息集合
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
