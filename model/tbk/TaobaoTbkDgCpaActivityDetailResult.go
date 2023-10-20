package tbk

// TaobaotbkdgcpaactivitydetailResult 结构体
type TaobaotbkdgcpaactivitydetailResult struct {
	// 错误代码
	Bizerrorfeature string `json:"biz_error_feature,omitempty" xml:"biz_error_feature,omitempty"`
	// 错误描述
	Bizerrordesc string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
	// 结果信息
	Resultmsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回素材id
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 结果码
	Resultcode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误代码
	Bizerrorcode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
