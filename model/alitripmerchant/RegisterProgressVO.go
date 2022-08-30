package alitripmerchant

// RegisterProgressVO 结构体
type RegisterProgressVO struct {
	// 在雅高官网注册时的手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 在雅高官网注册时的邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 是否需要弹验证码弹框
	PopUpVerification bool `json:"pop_up_verification,omitempty" xml:"pop_up_verification,omitempty"`
	// 注册状态
	RegisterStatus bool `json:"register_status,omitempty" xml:"register_status,omitempty"`
}
