package einvoice

import (
	"sync"
)

// InvoiceFlowRefundDto 结构体
type InvoiceFlowRefundDto struct {
	// 工单id(入驻、续约、加盘)
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// 业务平台码
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 退款备注信息
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
}

var poolInvoiceFlowRefundDto = sync.Pool{
	New: func() any {
		return new(InvoiceFlowRefundDto)
	},
}

// GetInvoiceFlowRefundDto() 从对象池中获取InvoiceFlowRefundDto
func GetInvoiceFlowRefundDto() *InvoiceFlowRefundDto {
	return poolInvoiceFlowRefundDto.Get().(*InvoiceFlowRefundDto)
}

// ReleaseInvoiceFlowRefundDto 释放InvoiceFlowRefundDto
func ReleaseInvoiceFlowRefundDto(v *InvoiceFlowRefundDto) {
	v.FlowId = ""
	v.PlatformCode = ""
	v.Remark = ""
	poolInvoiceFlowRefundDto.Put(v)
}
