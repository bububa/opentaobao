package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest 保险-开票结果反馈 API请求
// taobao.servindustry.finance.insurance.invoice.feedback
//
// 保险-开票结果反馈
type TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest struct {
	model.Params
	// 开票结果反馈请求
	_invoiceResultFeedbackTopReq *InvoiceResultFeedbackTopReq
}

// NewTaobaoservindustryfinanceinsuranceinvoicefeedbackRequest 初始化TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest对象
func NewTaobaoservindustryfinanceinsuranceinvoicefeedbackRequest() *TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest {
	return &TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.insurance.invoice.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceResultFeedbackTopReq is InvoiceResultFeedbackTopReq Setter
// 开票结果反馈请求
func (r *TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest) SetInvoiceResultFeedbackTopReq(_invoiceResultFeedbackTopReq *InvoiceResultFeedbackTopReq) error {
	r._invoiceResultFeedbackTopReq = _invoiceResultFeedbackTopReq
	r.Set("invoice_result_feedback_top_req", _invoiceResultFeedbackTopReq)
	return nil
}

// GetInvoiceResultFeedbackTopReq InvoiceResultFeedbackTopReq Getter
func (r TaobaoservindustryfinanceinsuranceinvoicefeedbackAPIRequest) GetInvoiceResultFeedbackTopReq() *InvoiceResultFeedbackTopReq {
	return r._invoiceResultFeedbackTopReq
}
