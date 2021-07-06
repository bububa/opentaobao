package wdk

// AlibabaWdkMerchantBrandQueryResult 结构体
type AlibabaWdkMerchantBrandQueryResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// result
	Results string `json:"results,omitempty" xml:"results,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
