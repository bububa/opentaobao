package drugtrace

// AlibabaAlihealthDrugKytSaveentResultModel 结构体
type AlibabaAlihealthDrugKytSaveentResultModel struct {
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 具体返回值
	Model *AlibabaAlihealthDrugKytSaveentModel `json:"model,omitempty" xml:"model,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
