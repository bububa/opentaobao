package einvoice

// ServiceResult 结构体
type ServiceResult struct {
	// 操作结果对象
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 系统自动生成
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 订购单信息
	Result *InvoiceFwOrderDto `json:"result,omitempty" xml:"result,omitempty"`
	// 续约返回结果
	InvoiceFlowRenewResult *InvoiceFlowRenewResultDto `json:"invoice_flow_renew_result,omitempty" xml:"invoice_flow_renew_result,omitempty"`
	// 工单详情
	TaxFlowQueryResult *TaxFlowQueryResult `json:"tax_flow_query_result,omitempty" xml:"tax_flow_query_result,omitempty"`
	// 申请结果
	ApplyResultDto *InvoiceApplyResultDto `json:"apply_result_dto,omitempty" xml:"apply_result_dto,omitempty"`
	// 发票申请详情
	InvoiceApplyDtl *InvoiceApplyDtlDto `json:"invoice_apply_dtl,omitempty" xml:"invoice_apply_dtl,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
