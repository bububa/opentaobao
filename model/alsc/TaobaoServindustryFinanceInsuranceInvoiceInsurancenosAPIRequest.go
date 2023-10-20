package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest 商家查询本次开票的保险单号 API请求
// taobao.servindustry.finance.insurance.invoice.insurancenos
//
// 商家查询本次开票的保险单号
type TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest struct {
	model.Params
	// 请求原文
	_invoiceInsuranceNoTopReq *InvoiceInsuranceNoTopReq
}

// NewTaobaoServindustryFinanceInsuranceInvoiceInsurancenosRequest 初始化TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest对象
func NewTaobaoServindustryFinanceInsuranceInvoiceInsurancenosRequest() *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest {
	return &TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest) Reset() {
	r._invoiceInsuranceNoTopReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.insurance.invoice.insurancenos"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceInsuranceNoTopReq is InvoiceInsuranceNoTopReq Setter
// 请求原文
func (r *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest) SetInvoiceInsuranceNoTopReq(_invoiceInsuranceNoTopReq *InvoiceInsuranceNoTopReq) error {
	r._invoiceInsuranceNoTopReq = _invoiceInsuranceNoTopReq
	r.Set("invoice_insurance_no_top_req", _invoiceInsuranceNoTopReq)
	return nil
}

// GetInvoiceInsuranceNoTopReq InvoiceInsuranceNoTopReq Getter
func (r TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest) GetInvoiceInsuranceNoTopReq() *InvoiceInsuranceNoTopReq {
	return r._invoiceInsuranceNoTopReq
}

var poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoServindustryFinanceInsuranceInvoiceInsurancenosRequest()
	},
}

// GetTaobaoServindustryFinanceInsuranceInvoiceInsurancenosRequest 从 sync.Pool 获取 TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest
func GetTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest() *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest {
	return poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest.Get().(*TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest)
}

// ReleaseTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest 将 TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest 放入 sync.Pool
func ReleaseTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest(v *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest) {
	v.Reset()
	poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosAPIRequest.Put(v)
}
