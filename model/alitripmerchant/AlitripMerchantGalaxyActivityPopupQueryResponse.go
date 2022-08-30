package alitripmerchant

// AlitripMerchantGalaxyActivityPopupQueryResponse 结构体
type AlitripMerchantGalaxyActivityPopupQueryResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 首页抽奖弹窗数据
	Content *ActivityDrawPopupVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
