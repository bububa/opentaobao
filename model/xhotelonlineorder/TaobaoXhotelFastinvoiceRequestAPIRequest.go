package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelFastinvoiceRequestAPIRequest 极速开票开票请求回传 API请求
// taobao.xhotel.fastinvoice.request
//
// 极速开票开票请求回传,用于记录航信开票请求数据
type TaobaoXhotelFastinvoiceRequestAPIRequest struct {
	model.Params
	// 无
	_invoiceInfoParam *InvoiceInfoParam
}

// NewTaobaoXhotelFastinvoiceRequestRequest 初始化TaobaoXhotelFastinvoiceRequestAPIRequest对象
func NewTaobaoXhotelFastinvoiceRequestRequest() *TaobaoXhotelFastinvoiceRequestAPIRequest {
	return &TaobaoXhotelFastinvoiceRequestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelFastinvoiceRequestAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.fastinvoice.request"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelFastinvoiceRequestAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetInvoiceInfoParam is InvoiceInfoParam Setter
// 无
func (r *TaobaoXhotelFastinvoiceRequestAPIRequest) SetInvoiceInfoParam(_invoiceInfoParam *InvoiceInfoParam) error {
	r._invoiceInfoParam = _invoiceInfoParam
	r.Set("invoice_info_param", _invoiceInfoParam)
	return nil
}

// GetInvoiceInfoParam InvoiceInfoParam Getter
func (r TaobaoXhotelFastinvoiceRequestAPIRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
	return r._invoiceInfoParam
}
