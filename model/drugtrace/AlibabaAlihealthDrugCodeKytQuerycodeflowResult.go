package drugtrace

// AlibabaAlihealthDrugCodeKytQuerycodeflowResult 结构体
type AlibabaAlihealthDrugCodeKytQuerycodeflowResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *AlibabaAlihealthDrugCodeKytQuerycodeflowModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
