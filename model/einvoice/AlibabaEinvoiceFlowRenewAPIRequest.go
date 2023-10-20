package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceFlowRenewAPIRequest 工单(入驻、加盘、续约)续约 API请求
// alibaba.einvoice.flow.renew
//
// 工单(含入驻、加盘、续约工单)续约能力开放
type AlibabaEinvoiceFlowRenewAPIRequest struct {
	model.Params
	// 续约请求参数
	_invoiceFlowRenewDto *InvoiceFlowRenewDto
}

// NewAlibabaEinvoiceFlowRenewRequest 初始化AlibabaEinvoiceFlowRenewAPIRequest对象
func NewAlibabaEinvoiceFlowRenewRequest() *AlibabaEinvoiceFlowRenewAPIRequest {
	return &AlibabaEinvoiceFlowRenewAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceFlowRenewAPIRequest) Reset() {
	r._invoiceFlowRenewDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceFlowRenewAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.flow.renew"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceFlowRenewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceFlowRenewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceFlowRenewDto is InvoiceFlowRenewDto Setter
// 续约请求参数
func (r *AlibabaEinvoiceFlowRenewAPIRequest) SetInvoiceFlowRenewDto(_invoiceFlowRenewDto *InvoiceFlowRenewDto) error {
	r._invoiceFlowRenewDto = _invoiceFlowRenewDto
	r.Set("invoice_flow_renew_dto", _invoiceFlowRenewDto)
	return nil
}

// GetInvoiceFlowRenewDto InvoiceFlowRenewDto Getter
func (r AlibabaEinvoiceFlowRenewAPIRequest) GetInvoiceFlowRenewDto() *InvoiceFlowRenewDto {
	return r._invoiceFlowRenewDto
}

var poolAlibabaEinvoiceFlowRenewAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceFlowRenewRequest()
	},
}

// GetAlibabaEinvoiceFlowRenewRequest 从 sync.Pool 获取 AlibabaEinvoiceFlowRenewAPIRequest
func GetAlibabaEinvoiceFlowRenewAPIRequest() *AlibabaEinvoiceFlowRenewAPIRequest {
	return poolAlibabaEinvoiceFlowRenewAPIRequest.Get().(*AlibabaEinvoiceFlowRenewAPIRequest)
}

// ReleaseAlibabaEinvoiceFlowRenewAPIRequest 将 AlibabaEinvoiceFlowRenewAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceFlowRenewAPIRequest(v *AlibabaEinvoiceFlowRenewAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceFlowRenewAPIRequest.Put(v)
}
