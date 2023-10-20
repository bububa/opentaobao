package einvoice

import (
	"sync"
)

// InvoiceOrderRefundResultDto 结构体
type InvoiceOrderRefundResultDto struct {
	// 拒绝退款原因，拒绝退款时必传
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 退款工单事件：  refund_agree: 服务商同意退款,  refund_reject: 服务商拒绝退单
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 退款工单流程ID
	FlowId string `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
}

var poolInvoiceOrderRefundResultDto = sync.Pool{
	New: func() any {
		return new(InvoiceOrderRefundResultDto)
	},
}

// GetInvoiceOrderRefundResultDto() 从对象池中获取InvoiceOrderRefundResultDto
func GetInvoiceOrderRefundResultDto() *InvoiceOrderRefundResultDto {
	return poolInvoiceOrderRefundResultDto.Get().(*InvoiceOrderRefundResultDto)
}

// ReleaseInvoiceOrderRefundResultDto 释放InvoiceOrderRefundResultDto
func ReleaseInvoiceOrderRefundResultDto(v *InvoiceOrderRefundResultDto) {
	v.Reason = ""
	v.Action = ""
	v.FlowId = ""
	poolInvoiceOrderRefundResultDto.Put(v)
}
