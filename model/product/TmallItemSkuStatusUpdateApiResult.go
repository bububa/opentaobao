package product

// TmallitemskustatusupdateApiResult 结构体
type TmallitemskustatusupdateApiResult struct {
	// 错误码集合，如有
	ErrorCodes []ErrorCode `json:"error_codes,omitempty" xml:"error_codes>error_code,omitempty"`
	// 执行结果信息
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
