package einvoice

// InvoiceFlowRenewResultDto 结构体
type InvoiceFlowRenewResultDto struct {
	// 续约工单ID
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 续约单结束时间
	ServEndTime string `json:"serv_end_time,omitempty" xml:"serv_end_time,omitempty"`
	// 续约单开始时间
	ServStartTime string `json:"serv_start_time,omitempty" xml:"serv_start_time,omitempty"`
}
