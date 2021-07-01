package jym

// RealNameVerifyTopDto 结构体
type RealNameVerifyTopDto struct {
	// 实名校验结果信息
	VerifyMsg string `json:"verify_msg,omitempty" xml:"verify_msg,omitempty"`
	// 实名校验结果信息
	VerifyCode int64 `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
	// 实名校验结果信息
	IdentityCode string `json:"identity_code,omitempty" xml:"identity_code,omitempty"`
}
