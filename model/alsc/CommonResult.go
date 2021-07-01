package alsc

// CommonResult 结构体
type CommonResult struct {
	// 成功状态
	BizSuccess bool `json:"biz_success,omitempty" xml:"biz_success,omitempty"`
	// 开卡是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 结果描述
	ResultDesc string `json:"result_desc,omitempty" xml:"result_desc,omitempty"`
	// 错误结果显示
	ResultView string `json:"result_view,omitempty" xml:"result_view,omitempty"`
}
