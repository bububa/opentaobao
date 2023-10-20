package alihouse

// TmallalihousetradecouponordercontractcredentialsqueryResult 结构体
type TmallalihousetradecouponordercontractcredentialsqueryResult struct {
	// 结果信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 结果code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 凭证对象
	Data *StsCredentialsDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
