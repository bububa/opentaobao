package drugtrace

// AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel 结构体
type AlibabaAlihealthDrugtraceTopYljgServiceGetenddateResultModel struct {
	// 服务截止时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
