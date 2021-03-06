package alihealthcrm

// AlibabaAlihealthPregnancyBaseinfoGetResult 结构体
type AlibabaAlihealthPregnancyBaseinfoGetResult struct {
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 数据集
	Model *AlibabaAlihealthPregnancyBaseinfoGetModel `json:"model,omitempty" xml:"model,omitempty"`
}
