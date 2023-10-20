package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusinvoicereturnAPIRequest 发票回调接口 API请求
// taobao.bus.invoice.return
//
// 汽车票发票回调接口
type TaobaobusinvoicereturnAPIRequest struct {
	model.Params
	// 入参对象
	_invoiceParam *ReceiptDo
}

// NewTaobaobusinvoicereturnRequest 初始化TaobaobusinvoicereturnAPIRequest对象
func NewTaobaobusinvoicereturnRequest() *TaobaobusinvoicereturnAPIRequest {
	return &TaobaobusinvoicereturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusinvoicereturnAPIRequest) GetApiMethodName() string {
	return "taobao.bus.invoice.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusinvoicereturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusinvoicereturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceParam is InvoiceParam Setter
// 入参对象
func (r *TaobaobusinvoicereturnAPIRequest) SetInvoiceParam(_invoiceParam *ReceiptDo) error {
	r._invoiceParam = _invoiceParam
	r.Set("invoice_param", _invoiceParam)
	return nil
}

// GetInvoiceParam InvoiceParam Getter
func (r TaobaobusinvoicereturnAPIRequest) GetInvoiceParam() *ReceiptDo {
	return r._invoiceParam
}
