package wdk

// OctopusOpenResult 结构体
type OctopusOpenResult struct {
	// 部分失败的商品编码列表
	FailedSkuCodes []string `json:"failed_sku_codes,omitempty" xml:"failed_sku_codes>string,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 活动ID
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 操作是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
