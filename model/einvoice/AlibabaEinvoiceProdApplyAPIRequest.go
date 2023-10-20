package einvoice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceProdApplyAPIRequest 提交发票申请 API请求
// alibaba.einvoice.prod.apply
//
// 提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
type AlibabaEinvoiceProdApplyAPIRequest struct {
	model.Params
	// 申请开票请求
	_paramInvoiceApplyDto *InvoiceApplyDto
}

// NewAlibabaEinvoiceProdApplyRequest 初始化AlibabaEinvoiceProdApplyAPIRequest对象
func NewAlibabaEinvoiceProdApplyRequest() *AlibabaEinvoiceProdApplyAPIRequest {
	return &AlibabaEinvoiceProdApplyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaEinvoiceProdApplyAPIRequest) Reset() {
	r._paramInvoiceApplyDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceProdApplyAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.prod.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceProdApplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaEinvoiceProdApplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamInvoiceApplyDto is ParamInvoiceApplyDto Setter
// 申请开票请求
func (r *AlibabaEinvoiceProdApplyAPIRequest) SetParamInvoiceApplyDto(_paramInvoiceApplyDto *InvoiceApplyDto) error {
	r._paramInvoiceApplyDto = _paramInvoiceApplyDto
	r.Set("param_invoice_apply_dto", _paramInvoiceApplyDto)
	return nil
}

// GetParamInvoiceApplyDto ParamInvoiceApplyDto Getter
func (r AlibabaEinvoiceProdApplyAPIRequest) GetParamInvoiceApplyDto() *InvoiceApplyDto {
	return r._paramInvoiceApplyDto
}

var poolAlibabaEinvoiceProdApplyAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaEinvoiceProdApplyRequest()
	},
}

// GetAlibabaEinvoiceProdApplyRequest 从 sync.Pool 获取 AlibabaEinvoiceProdApplyAPIRequest
func GetAlibabaEinvoiceProdApplyAPIRequest() *AlibabaEinvoiceProdApplyAPIRequest {
	return poolAlibabaEinvoiceProdApplyAPIRequest.Get().(*AlibabaEinvoiceProdApplyAPIRequest)
}

// ReleaseAlibabaEinvoiceProdApplyAPIRequest 将 AlibabaEinvoiceProdApplyAPIRequest 放入 sync.Pool
func ReleaseAlibabaEinvoiceProdApplyAPIRequest(v *AlibabaEinvoiceProdApplyAPIRequest) {
	v.Reset()
	poolAlibabaEinvoiceProdApplyAPIRequest.Put(v)
}
