package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelFastinvoiceCompleteAPIRequest 极速开票开票请求完成 API请求
// taobao.xhotel.fastinvoice.complete
//
// 极速开票开票请求回传,用于更新航信开票请求数据
type TaobaoXhotelFastinvoiceCompleteAPIRequest struct {
	model.Params
	// 无
	_invoiceInfoParam *InvoiceInfoParam
}

// NewTaobaoXhotelFastinvoiceCompleteRequest 初始化TaobaoXhotelFastinvoiceCompleteAPIRequest对象
func NewTaobaoXhotelFastinvoiceCompleteRequest() *TaobaoXhotelFastinvoiceCompleteAPIRequest {
	return &TaobaoXhotelFastinvoiceCompleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelFastinvoiceCompleteAPIRequest) Reset() {
	r._invoiceInfoParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelFastinvoiceCompleteAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.fastinvoice.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelFastinvoiceCompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelFastinvoiceCompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceInfoParam is InvoiceInfoParam Setter
// 无
func (r *TaobaoXhotelFastinvoiceCompleteAPIRequest) SetInvoiceInfoParam(_invoiceInfoParam *InvoiceInfoParam) error {
	r._invoiceInfoParam = _invoiceInfoParam
	r.Set("invoice_info_param", _invoiceInfoParam)
	return nil
}

// GetInvoiceInfoParam InvoiceInfoParam Getter
func (r TaobaoXhotelFastinvoiceCompleteAPIRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
	return r._invoiceInfoParam
}

var poolTaobaoXhotelFastinvoiceCompleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelFastinvoiceCompleteRequest()
	},
}

// GetTaobaoXhotelFastinvoiceCompleteRequest 从 sync.Pool 获取 TaobaoXhotelFastinvoiceCompleteAPIRequest
func GetTaobaoXhotelFastinvoiceCompleteAPIRequest() *TaobaoXhotelFastinvoiceCompleteAPIRequest {
	return poolTaobaoXhotelFastinvoiceCompleteAPIRequest.Get().(*TaobaoXhotelFastinvoiceCompleteAPIRequest)
}

// ReleaseTaobaoXhotelFastinvoiceCompleteAPIRequest 将 TaobaoXhotelFastinvoiceCompleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelFastinvoiceCompleteAPIRequest(v *TaobaoXhotelFastinvoiceCompleteAPIRequest) {
	v.Reset()
	poolTaobaoXhotelFastinvoiceCompleteAPIRequest.Put(v)
}
