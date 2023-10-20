package alihealth2

// AlibabaalihealthdentalstoreinsertorupdateMtopResult 结构体
type AlibabaalihealthdentalstoreinsertorupdateMtopResult struct {
	// msg_code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msg_info
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *StoreAuditVo `json:"model,omitempty" xml:"model,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
