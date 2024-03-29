package einvoice

import (
	"sync"
)

// InvoiceFlowRenewDto 结构体
type InvoiceFlowRenewDto struct {
	// 外部幂等ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 被续约工单ID
	ParentFlowId string `json:"parent_flow_id,omitempty" xml:"parent_flow_id,omitempty"`
	// 请求来源平台code,由中台生成
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 服务的有效天数，单位为天
	ServiceValidDays int64 `json:"service_valid_days,omitempty" xml:"service_valid_days,omitempty"`
}

var poolInvoiceFlowRenewDto = sync.Pool{
	New: func() any {
		return new(InvoiceFlowRenewDto)
	},
}

// GetInvoiceFlowRenewDto() 从对象池中获取InvoiceFlowRenewDto
func GetInvoiceFlowRenewDto() *InvoiceFlowRenewDto {
	return poolInvoiceFlowRenewDto.Get().(*InvoiceFlowRenewDto)
}

// ReleaseInvoiceFlowRenewDto 释放InvoiceFlowRenewDto
func ReleaseInvoiceFlowRenewDto(v *InvoiceFlowRenewDto) {
	v.OuterId = ""
	v.ParentFlowId = ""
	v.PlatformCode = ""
	v.ServiceValidDays = 0
	poolInvoiceFlowRenewDto.Put(v)
}
