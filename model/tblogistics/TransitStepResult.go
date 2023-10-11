package tblogistics

// TransitStepResult 结构体
type TransitStepResult struct {
	// 列表
	TraceList []TransitStepInfo `json:"trace_list,omitempty" xml:"trace_list>transit_step_info,omitempty"`
	// 运单号
	OutSid string `json:"out_sid,omitempty" xml:"out_sid,omitempty"`
	// 物流公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 订单物流状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 交易号
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}
