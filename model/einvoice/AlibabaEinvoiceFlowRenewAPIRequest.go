package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceFlowRenewAPIRequest
工单(入驻、加盘、续约)续约 API请求
alibaba.einvoice.flow.renew

工单(含入驻、加盘、续约工单)续约能力开放 */
type AlibabaEinvoiceFlowRenewAPIRequest struct {
	model.Params
	// 续约请求参数
	_invoiceFlowRenewDto *InvoiceFlowRenewDto
}

// NewAlibabaEinvoiceFlowRenewRequest 初始化AlibabaEinvoiceFlowRenewAPIRequest对象
func NewAlibabaEinvoiceFlowRenewRequest() *AlibabaEinvoiceFlowRenewAPIRequest {
	return &AlibabaEinvoiceFlowRenewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowRenewAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.flow.renew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowRenewAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InvoiceFlowRenewDto Setter
// 续约请求参数
func (r *AlibabaEinvoiceFlowRenewAPIRequest) SetInvoiceFlowRenewDto(_invoiceFlowRenewDto *InvoiceFlowRenewDto) error {
	r._invoiceFlowRenewDto = _invoiceFlowRenewDto
	r.Set("invoice_flow_renew_dto", _invoiceFlowRenewDto)
	return nil
}

// Get InvoiceFlowRenewDto Getter
func (r AlibabaEinvoiceFlowRenewAPIRequest) GetInvoiceFlowRenewDto() *InvoiceFlowRenewDto {
	return r._invoiceFlowRenewDto
}
