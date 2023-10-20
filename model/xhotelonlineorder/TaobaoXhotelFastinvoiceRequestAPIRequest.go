package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelfastinvoicerequestAPIRequest 极速开票开票请求回传 API请求
// taobao.xhotel.fastinvoice.request
//
// 极速开票开票请求回传,用于记录航信开票请求数据
type TaobaoxhotelfastinvoicerequestAPIRequest struct {
	model.Params
	// 无
	_invoiceInfoParam *InvoiceInfoParam
}

// NewTaobaoxhotelfastinvoicerequestRequest 初始化TaobaoxhotelfastinvoicerequestAPIRequest对象
func NewTaobaoxhotelfastinvoicerequestRequest() *TaobaoxhotelfastinvoicerequestAPIRequest {
	return &TaobaoxhotelfastinvoicerequestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelfastinvoicerequestAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.fastinvoice.request"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelfastinvoicerequestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelfastinvoicerequestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceInfoParam is InvoiceInfoParam Setter
// 无
func (r *TaobaoxhotelfastinvoicerequestAPIRequest) SetInvoiceInfoParam(_invoiceInfoParam *InvoiceInfoParam) error {
	r._invoiceInfoParam = _invoiceInfoParam
	r.Set("invoice_info_param", _invoiceInfoParam)
	return nil
}

// GetInvoiceInfoParam InvoiceInfoParam Getter
func (r TaobaoxhotelfastinvoicerequestAPIRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
	return r._invoiceInfoParam
}
