package tbk

// TaobaoTbkDgCpaActivityDetailResult 结构体
type TaobaoTbkDgCpaActivityDetailResult struct {
	// 错误代码
	BizErrorFeature string `json:"biz_error_feature,omitempty" xml:"biz_error_feature,omitempty"`
	// 错误描述
	BizErrorDesc string `json:"biz_error_desc,omitempty" xml:"biz_error_desc,omitempty"`
	// 结果信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回素材id
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 结果码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误代码
	BizErrorCode int64 `json:"biz_error_code,omitempty" xml:"biz_error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
