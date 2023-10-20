package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest 保险-开票结果反馈 API请求
// taobao.servindustry.finance.insurance.invoice.feedback
//
// 保险-开票结果反馈
type TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest struct {
	model.Params
	// 开票结果反馈请求
	_invoiceResultFeedbackTopReq *InvoiceResultFeedbackTopReq
}

// NewTaobaoServindustryFinanceInsuranceInvoiceFeedbackRequest 初始化TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest对象
func NewTaobaoServindustryFinanceInsuranceInvoiceFeedbackRequest() *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest {
	return &TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest) Reset() {
	r._invoiceResultFeedbackTopReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.insurance.invoice.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceResultFeedbackTopReq is InvoiceResultFeedbackTopReq Setter
// 开票结果反馈请求
func (r *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest) SetInvoiceResultFeedbackTopReq(_invoiceResultFeedbackTopReq *InvoiceResultFeedbackTopReq) error {
	r._invoiceResultFeedbackTopReq = _invoiceResultFeedbackTopReq
	r.Set("invoice_result_feedback_top_req", _invoiceResultFeedbackTopReq)
	return nil
}

// GetInvoiceResultFeedbackTopReq InvoiceResultFeedbackTopReq Getter
func (r TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest) GetInvoiceResultFeedbackTopReq() *InvoiceResultFeedbackTopReq {
	return r._invoiceResultFeedbackTopReq
}

var poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoServindustryFinanceInsuranceInvoiceFeedbackRequest()
	},
}

// GetTaobaoServindustryFinanceInsuranceInvoiceFeedbackRequest 从 sync.Pool 获取 TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest
func GetTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest() *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest {
	return poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest.Get().(*TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest)
}

// ReleaseTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest 将 TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest 放入 sync.Pool
func ReleaseTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest(v *TaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest) {
	v.Reset()
	poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackAPIRequest.Put(v)
}
