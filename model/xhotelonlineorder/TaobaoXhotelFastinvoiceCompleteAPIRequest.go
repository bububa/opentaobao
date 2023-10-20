package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelfastinvoicecompleteAPIRequest 极速开票开票请求完成 API请求
// taobao.xhotel.fastinvoice.complete
//
// 极速开票开票请求回传,用于更新航信开票请求数据
type TaobaoxhotelfastinvoicecompleteAPIRequest struct {
	model.Params
	// 无
	_invoiceInfoParam *InvoiceInfoParam
}

// NewTaobaoxhotelfastinvoicecompleteRequest 初始化TaobaoxhotelfastinvoicecompleteAPIRequest对象
func NewTaobaoxhotelfastinvoicecompleteRequest() *TaobaoxhotelfastinvoicecompleteAPIRequest {
	return &TaobaoxhotelfastinvoicecompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelfastinvoicecompleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.fastinvoice.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelfastinvoicecompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelfastinvoicecompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceInfoParam is InvoiceInfoParam Setter
// 无
func (r *TaobaoxhotelfastinvoicecompleteAPIRequest) SetInvoiceInfoParam(_invoiceInfoParam *InvoiceInfoParam) error {
	r._invoiceInfoParam = _invoiceInfoParam
	r.Set("invoice_info_param", _invoiceInfoParam)
	return nil
}

// GetInvoiceInfoParam InvoiceInfoParam Getter
func (r TaobaoxhotelfastinvoicecompleteAPIRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
	return r._invoiceInfoParam
}
