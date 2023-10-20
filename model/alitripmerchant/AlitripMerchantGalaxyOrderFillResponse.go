package alitripmerchant

// AlitripmerchantgalaxyorderfillResponse 结构体
type AlitripmerchantgalaxyorderfillResponse struct {
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 填单页展示对象
	Content *FillPageVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
