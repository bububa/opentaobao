package fpm

import (
	"sync"
)

// InvoiceRegisterRequest 结构体
type InvoiceRegisterRequest struct {
	// 操作人
	OperateBy string `json:"operate_by,omitempty" xml:"operate_by,omitempty"`
	// 业务请求号【必填】
	RequestNo string `json:"request_no,omitempty" xml:"request_no,omitempty"`
	// 业务平台代码【必填】
	PlatformCode string `json:"platform_code,omitempty" xml:"platform_code,omitempty"`
	// 发票信息
	InvoiceDTO *RegisterInvoiceDto `json:"invoice_d_t_o,omitempty" xml:"invoice_d_t_o,omitempty"`
}

var poolInvoiceRegisterRequest = sync.Pool{
	New: func() any {
		return new(InvoiceRegisterRequest)
	},
}

// GetInvoiceRegisterRequest() 从对象池中获取InvoiceRegisterRequest
func GetInvoiceRegisterRequest() *InvoiceRegisterRequest {
	return poolInvoiceRegisterRequest.Get().(*InvoiceRegisterRequest)
}

// ReleaseInvoiceRegisterRequest 释放InvoiceRegisterRequest
func ReleaseInvoiceRegisterRequest(v *InvoiceRegisterRequest) {
	v.OperateBy = ""
	v.RequestNo = ""
	v.PlatformCode = ""
	v.InvoiceDTO = nil
	poolInvoiceRegisterRequest.Put(v)
}
