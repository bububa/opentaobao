package tbk

// TaobaotbkdgvegastljcreateResult 结构体
type TaobaotbkdgvegastljcreateResult struct {
	// msgCode
	Msgcode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	Msginfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *RightsInstanceCreateResult `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
