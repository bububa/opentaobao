package bus

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusInvoiceReturnAPIRequest) Reset() {
	r._invoiceParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusInvoiceReturnAPIRequest) GetApiMethodName() string {
	return "taobao.bus.invoice.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusInvoiceReturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusInvoiceReturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceParam is InvoiceParam Setter
// 入参对象
func (r *TaobaoBusInvoiceReturnAPIRequest) SetInvoiceParam(_invoiceParam *ReceiptDo) error {
	r._invoiceParam = _invoiceParam
	r.Set("invoice_param", _invoiceParam)
	return nil
}

// GetInvoiceParam InvoiceParam Getter
func (r TaobaoBusInvoiceReturnAPIRequest) GetInvoiceParam() *ReceiptDo {
	return r._invoiceParam
}

var poolTaobaoBusInvoiceReturnAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusInvoiceReturnRequest()
	},
}

// GetTaobaoBusInvoiceReturnRequest 从 sync.Pool 获取 TaobaoBusInvoiceReturnAPIRequest
func GetTaobaoBusInvoiceReturnAPIRequest() *TaobaoBusInvoiceReturnAPIRequest {
	return poolTaobaoBusInvoiceReturnAPIRequest.Get().(*TaobaoBusInvoiceReturnAPIRequest)
}

// ReleaseTaobaoBusInvoiceReturnAPIRequest 将 TaobaoBusInvoiceReturnAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusInvoiceReturnAPIRequest(v *TaobaoBusInvoiceReturnAPIRequest) {
	v.Reset()
	poolTaobaoBusInvoiceReturnAPIRequest.Put(v)
}
