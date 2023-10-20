package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceprodapplyAPIRequest 提交发票申请 API请求
// alibaba.einvoice.prod.apply
//
// 提交开票申请，如果商户授权自动开票则自动转开票，否则等待商户审核。
type AlibabaeinvoiceprodapplyAPIRequest struct {
	model.Params
	// 申请开票请求
	_paramInvoiceApplyDto *InvoiceApplyDto
}

// NewAlibabaeinvoiceprodapplyRequest 初始化AlibabaeinvoiceprodapplyAPIRequest对象
func NewAlibabaeinvoiceprodapplyRequest() *AlibabaeinvoiceprodapplyAPIRequest {
	return &AlibabaeinvoiceprodapplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceprodapplyAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.prod.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceprodapplyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceprodapplyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamInvoiceApplyDto is ParamInvoiceApplyDto Setter
// 申请开票请求
func (r *AlibabaeinvoiceprodapplyAPIRequest) SetParamInvoiceApplyDto(_paramInvoiceApplyDto *InvoiceApplyDto) error {
	r._paramInvoiceApplyDto = _paramInvoiceApplyDto
	r.Set("param_invoice_apply_dto", _paramInvoiceApplyDto)
	return nil
}

// GetParamInvoiceApplyDto ParamInvoiceApplyDto Getter
func (r AlibabaeinvoiceprodapplyAPIRequest) GetParamInvoiceApplyDto() *InvoiceApplyDto {
	return r._paramInvoiceApplyDto
}
