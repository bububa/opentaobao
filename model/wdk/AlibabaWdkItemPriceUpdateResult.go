package wdk

// AlibabaWdkItemPriceUpdateResult 结构体
type AlibabaWdkItemPriceUpdateResult struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
}
