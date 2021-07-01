package alihealthoutflow

// AlibabaAlihealthHealthRecordHaveResult 结构体
type AlibabaAlihealthHealthRecordHaveResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 返回结果
	Model *HaveRecordResponse `json:"model,omitempty" xml:"model,omitempty"`
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
