package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceflowrefundAPIRequest 退订工单(入驻、加盘、续约) API请求
// alibaba.einvoice.flow.refund
//
// 电子发票工单系统，工单退订能力开放
type AlibabaeinvoiceflowrefundAPIRequest struct {
	model.Params
	// 退订请求参数
	_invoiceFlowRefund *InvoiceFlowRefundDto
}

// NewAlibabaeinvoiceflowrefundRequest 初始化AlibabaeinvoiceflowrefundAPIRequest对象
func NewAlibabaeinvoiceflowrefundRequest() *AlibabaeinvoiceflowrefundAPIRequest {
	return &AlibabaeinvoiceflowrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceflowrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.flow.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceflowrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceflowrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceFlowRefund is InvoiceFlowRefund Setter
// 退订请求参数
func (r *AlibabaeinvoiceflowrefundAPIRequest) SetInvoiceFlowRefund(_invoiceFlowRefund *InvoiceFlowRefundDto) error {
	r._invoiceFlowRefund = _invoiceFlowRefund
	r.Set("invoice_flow_refund", _invoiceFlowRefund)
	return nil
}

// GetInvoiceFlowRefund InvoiceFlowRefund Getter
func (r AlibabaeinvoiceflowrefundAPIRequest) GetInvoiceFlowRefund() *InvoiceFlowRefundDto {
	return r._invoiceFlowRefund
}
