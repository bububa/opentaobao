package drugtrace

// AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgQueryGetbyrefentidResultModel struct {
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象
	Model *PUserEntInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否响应成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
