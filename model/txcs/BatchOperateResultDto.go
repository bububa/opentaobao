package txcs

// BatchOperateResultDto 结构体
type BatchOperateResultDto struct {
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 失败列表
	FailList []InvoiceInputResultDto `json:"fail_list,omitempty" xml:"fail_list>invoice_input_result_dto,omitempty"`
	// 成功列表
	SuccessList []InvoiceInputResultDto `json:"success_list,omitempty" xml:"success_list>invoice_input_result_dto,omitempty"`
}
