package drugtrace

// AlibabaAlihealthDrugKytYbGetcoderelationResult 结构体
type AlibabaAlihealthDrugKytYbGetcoderelationResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// model
	ModelList *AlibabaAlihealthDrugKytYbGetcoderelationModel `json:"model_list,omitempty" xml:"model_list,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
