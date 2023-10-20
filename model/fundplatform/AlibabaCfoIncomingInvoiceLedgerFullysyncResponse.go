package fundplatform

import (
	"sync"
)

// AlibabaCfoIncomingInvoiceLedgerFullysyncResponse 结构体
type AlibabaCfoIncomingInvoiceLedgerFullysyncResponse struct {
	// 响应码
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 响应消息
	ResponseMsg string `json:"response_msg,omitempty" xml:"response_msg,omitempty"`
	// 是否成功
	Succeeded string `json:"succeeded,omitempty" xml:"succeeded,omitempty"`
}

var poolAlibabaCfoIncomingInvoiceLedgerFullysyncResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCfoIncomingInvoiceLedgerFullysyncResponse)
	},
}

// GetAlibabaCfoIncomingInvoiceLedgerFullysyncResponse() 从对象池中获取AlibabaCfoIncomingInvoiceLedgerFullysyncResponse
func GetAlibabaCfoIncomingInvoiceLedgerFullysyncResponse() *AlibabaCfoIncomingInvoiceLedgerFullysyncResponse {
	return poolAlibabaCfoIncomingInvoiceLedgerFullysyncResponse.Get().(*AlibabaCfoIncomingInvoiceLedgerFullysyncResponse)
}

// ReleaseAlibabaCfoIncomingInvoiceLedgerFullysyncResponse 释放AlibabaCfoIncomingInvoiceLedgerFullysyncResponse
func ReleaseAlibabaCfoIncomingInvoiceLedgerFullysyncResponse(v *AlibabaCfoIncomingInvoiceLedgerFullysyncResponse) {
	v.ResponseCode = ""
	v.ResponseMsg = ""
	v.Succeeded = ""
	poolAlibabaCfoIncomingInvoiceLedgerFullysyncResponse.Put(v)
}
