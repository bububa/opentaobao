package einvoice

import (
	"sync"
)

// InvoiceFlowRenewResultDto 结构体
type InvoiceFlowRenewResultDto struct {
	// 续约工单ID
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 续约单结束时间
	ServEndTime string `json:"serv_end_time,omitempty" xml:"serv_end_time,omitempty"`
	// 续约单开始时间
	ServStartTime string `json:"serv_start_time,omitempty" xml:"serv_start_time,omitempty"`
}

var poolInvoiceFlowRenewResultDto = sync.Pool{
	New: func() any {
		return new(InvoiceFlowRenewResultDto)
	},
}

// GetInvoiceFlowRenewResultDto() 从对象池中获取InvoiceFlowRenewResultDto
func GetInvoiceFlowRenewResultDto() *InvoiceFlowRenewResultDto {
	return poolInvoiceFlowRenewResultDto.Get().(*InvoiceFlowRenewResultDto)
}

// ReleaseInvoiceFlowRenewResultDto 释放InvoiceFlowRenewResultDto
func ReleaseInvoiceFlowRenewResultDto(v *InvoiceFlowRenewResultDto) {
	v.FlowId = ""
	v.ServEndTime = ""
	v.ServStartTime = ""
	poolInvoiceFlowRenewResultDto.Put(v)
}
