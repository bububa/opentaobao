package cainiaohandover

// HandoverContentUpdateErrorParcelDto 结构体
type HandoverContentUpdateErrorParcelDto struct {
	// 小包LP号
	LpCode string `json:"lp_code,omitempty" xml:"lp_code,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误文案
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
