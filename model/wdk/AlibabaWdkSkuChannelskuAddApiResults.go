package wdk

// AlibabaWdkSkuChannelskuAddApiResults 结构体
type AlibabaWdkSkuChannelskuAddApiResults struct {
	// 返会结果集合
	Models []AlibabaWdkSkuChannelskuAddApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_channelsku_add_api_result,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
