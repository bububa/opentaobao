package xhotelonlineorder

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelFastinvoiceRequestAPIRequest) Reset() {
	r._invoiceInfoParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelFastinvoiceRequestAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.fastinvoice.request"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelFastinvoiceRequestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelFastinvoiceRequestAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoXhotelFastinvoiceRequestAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelFastinvoiceRequestRequest()
	},
}

// GetTaobaoXhotelFastinvoiceRequestRequest 从 sync.Pool 获取 TaobaoXhotelFastinvoiceRequestAPIRequest
func GetTaobaoXhotelFastinvoiceRequestAPIRequest() *TaobaoXhotelFastinvoiceRequestAPIRequest {
	return poolTaobaoXhotelFastinvoiceRequestAPIRequest.Get().(*TaobaoXhotelFastinvoiceRequestAPIRequest)
}

// ReleaseTaobaoXhotelFastinvoiceRequestAPIRequest 将 TaobaoXhotelFastinvoiceRequestAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelFastinvoiceRequestAPIRequest(v *TaobaoXhotelFastinvoiceRequestAPIRequest) {
	v.Reset()
	poolTaobaoXhotelFastinvoiceRequestAPIRequest.Put(v)
}
