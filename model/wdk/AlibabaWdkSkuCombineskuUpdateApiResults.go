package wdk

// AlibabaWdkSkuCombineskuUpdateApiResults 结构体
type AlibabaWdkSkuCombineskuUpdateApiResults struct {
	// 商品列表
	Models []AlibabaWdkSkuCombineskuUpdateApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_combinesku_update_api_result,omitempty"`
	// 接口调用异常编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 接口调用异常信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
