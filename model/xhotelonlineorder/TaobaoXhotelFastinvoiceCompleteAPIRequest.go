package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelFastinvoiceCompleteAPIRequest
极速开票开票请求完成 API请求
taobao.xhotel.fastinvoice.complete

极速开票开票请求回传,用于更新航信开票请求数据 */
type TaobaoXhotelFastinvoiceCompleteAPIRequest struct {
	model.Params
	// 无
	_invoiceInfoParam *InvoiceInfoParam
}

// NewTaobaoXhotelFastinvoiceCompleteRequest 初始化TaobaoXhotelFastinvoiceCompleteAPIRequest对象
func NewTaobaoXhotelFastinvoiceCompleteRequest() *TaobaoXhotelFastinvoiceCompleteAPIRequest {
	return &TaobaoXhotelFastinvoiceCompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelFastinvoiceCompleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.fastinvoice.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelFastinvoiceCompleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InvoiceInfoParam Setter
// 无
func (r *TaobaoXhotelFastinvoiceCompleteAPIRequest) SetInvoiceInfoParam(_invoiceInfoParam *InvoiceInfoParam) error {
	r._invoiceInfoParam = _invoiceInfoParam
	r.Set("invoice_info_param", _invoiceInfoParam)
	return nil
}

// Get InvoiceInfoParam Getter
func (r TaobaoXhotelFastinvoiceCompleteAPIRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
	return r._invoiceInfoParam
}
