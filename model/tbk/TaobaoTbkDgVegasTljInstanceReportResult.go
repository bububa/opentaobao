package tbk

// TaobaoTbkDgVegasTljInstanceReportResult 结构体
type TaobaoTbkDgVegasTljInstanceReportResult struct {
	// model
	Model *TljInstanceReportDto `json:"model,omitempty" xml:"model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
