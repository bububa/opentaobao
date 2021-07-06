package drugtrace

// AlibabaAlihealthDrugCodeKytYyApplycodeResult 结构体
type AlibabaAlihealthDrugCodeKytYyApplycodeResult struct {
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 父子码关系对象
	Model *AlibabaAlihealthDrugCodeKytYyApplycodeModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
