package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceflowrenewAPIRequest 工单(入驻、加盘、续约)续约 API请求
// alibaba.einvoice.flow.renew
//
// 工单(含入驻、加盘、续约工单)续约能力开放
type AlibabaeinvoiceflowrenewAPIRequest struct {
	model.Params
	// 续约请求参数
	_invoiceFlowRenewDto *InvoiceFlowRenewDto
}

// NewAlibabaeinvoiceflowrenewRequest 初始化AlibabaeinvoiceflowrenewAPIRequest对象
func NewAlibabaeinvoiceflowrenewRequest() *AlibabaeinvoiceflowrenewAPIRequest {
	return &AlibabaeinvoiceflowrenewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceflowrenewAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.flow.renew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceflowrenewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceflowrenewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceFlowRenewDto is InvoiceFlowRenewDto Setter
// 续约请求参数
func (r *AlibabaeinvoiceflowrenewAPIRequest) SetInvoiceFlowRenewDto(_invoiceFlowRenewDto *InvoiceFlowRenewDto) error {
	r._invoiceFlowRenewDto = _invoiceFlowRenewDto
	r.Set("invoice_flow_renew_dto", _invoiceFlowRenewDto)
	return nil
}

// GetInvoiceFlowRenewDto InvoiceFlowRenewDto Getter
func (r AlibabaeinvoiceflowrenewAPIRequest) GetInvoiceFlowRenewDto() *InvoiceFlowRenewDto {
	return r._invoiceFlowRenewDto
}
