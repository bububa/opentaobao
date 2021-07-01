package wdk

// AlibabaTclsAelophyMerchantChannelRefundCancelApiResult 结构体
type AlibabaTclsAelophyMerchantChannelRefundCancelApiResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回码说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
