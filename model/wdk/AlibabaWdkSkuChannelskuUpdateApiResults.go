package wdk

// AlibabaWdkSkuChannelskuUpdateApiResults 结构体
type AlibabaWdkSkuChannelskuUpdateApiResults struct {
	// 单个商品返回结果集合
	Models []AlibabaWdkSkuChannelskuUpdateApiResult `json:"models,omitempty" xml:"models>alibaba_wdk_sku_channelsku_update_api_result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
}
