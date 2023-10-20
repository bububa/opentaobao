package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceFlowRefundAPIRequest 退订工单(入驻、加盘、续约) API请求
// alibaba.einvoice.flow.refund
//
// 电子发票工单系统，工单退订能力开放
type AlibabaEinvoiceFlowRefundAPIRequest struct {
	model.Params
	// 退订请求参数
	_invoiceFlowRefund *InvoiceFlowRefundDto
}

// NewAlibabaEinvoiceFlowRefundRequest 初始化AlibabaEinvoiceFlowRefundAPIRequest对象
func NewAlibabaEinvoiceFlowRefundRequest() *AlibabaEinvoiceFlowRefundAPIRequest {
	return &AlibabaEinvoiceFlowRefundAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceFlowRefundAPIRequest) Reset() {
	r._invoiceFlowRefund = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.flow.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceFlowRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceFlowRefund is InvoiceFlowRefund Setter
// 退订请求参数
func (r *AlibabaEinvoiceFlowRefundAPIRequest) SetInvoiceFlowRefund(_invoiceFlowRefund *InvoiceFlowRefundDto) error {
	r._invoiceFlowRefund = _invoiceFlowRefund
	r.Set("invoice_flow_refund", _invoiceFlowRefund)
	return nil
}

// GetInvoiceFlowRefund InvoiceFlowRefund Getter
func (r AlibabaEinvoiceFlowRefundAPIRequest) GetInvoiceFlowRefund() *InvoiceFlowRefundDto {
	return r._invoiceFlowRefund
}

var poolAlibabaEinvoiceFlowRefundAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceFlowRefundRequest()
	},
}

// GetAlibabaEinvoiceFlowRefundRequest 从 sync.Pool 获取 AlibabaEinvoiceFlowRefundAPIRequest
func GetAlibabaEinvoiceFlowRefundAPIRequest() *AlibabaEinvoiceFlowRefundAPIRequest {
	return poolAlibabaEinvoiceFlowRefundAPIRequest.Get().(*AlibabaEinvoiceFlowRefundAPIRequest)
}

// ReleaseAlibabaEinvoiceFlowRefundAPIRequest 将 AlibabaEinvoiceFlowRefundAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceFlowRefundAPIRequest(v *AlibabaEinvoiceFlowRefundAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceFlowRefundAPIRequest.Put(v)
}
