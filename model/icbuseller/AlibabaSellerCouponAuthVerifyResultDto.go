package icbuseller

// AlibabaSellerCouponAuthVerifyResultDto 结构体
type AlibabaSellerCouponAuthVerifyResultDto struct {
	// 验证失败结果
	ExecDescription string `json:"exec_description,omitempty" xml:"exec_description,omitempty"`
	// 返回码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 是否正常返回
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否验证通过
	Dto bool `json:"dto,omitempty" xml:"dto,omitempty"`
}
