package wdk

// AlibabawdkmerchantitemqueryResult 结构体
type AlibabawdkmerchantitemqueryResult struct {
	// errorCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errorDesc
	ErrDesc string `json:"err_desc,omitempty" xml:"err_desc,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Suc bool `json:"suc,omitempty" xml:"suc,omitempty"`
}
