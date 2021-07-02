package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusInvoiceReturnAPIRequest 发票回调接口 API请求
// taobao.bus.invoice.return
//
// 汽车票发票回调接口
type TaobaoBusInvoiceReturnAPIRequest struct {
	model.Params
	// 入参对象
	_invoiceParam *ReceiptDo
}

// NewTaobaoBusInvoiceReturnRequest 初始化TaobaoBusInvoiceReturnAPIRequest对象
func NewTaobaoBusInvoiceReturnRequest() *TaobaoBusInvoiceReturnAPIRequest {
	return &TaobaoBusInvoiceReturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusInvoiceReturnAPIRequest) GetApiMethodName() string {
	return "taobao.bus.invoice.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusInvoiceReturnAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InvoiceParam Setter
// 入参对象
func (r *TaobaoBusInvoiceReturnAPIRequest) SetInvoiceParam(_invoiceParam *ReceiptDo) error {
	r._invoiceParam = _invoiceParam
	r.Set("invoice_param", _invoiceParam)
	return nil
}

// Get InvoiceParam Getter
func (r TaobaoBusInvoiceReturnAPIRequest) GetInvoiceParam() *ReceiptDo {
	return r._invoiceParam
}
