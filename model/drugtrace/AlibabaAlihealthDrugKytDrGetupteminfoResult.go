package drugtrace

// AlibabaAlihealthDrugKytDrGetupteminfoResult 结构体
type AlibabaAlihealthDrugKytDrGetupteminfoResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *AlibabaAlihealthDrugKytDrGetupteminfoModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
