package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomepaymentmethodsyncAPIRequest 付款方式上翻 API请求
// alibaba.alihouse.newhome.payment.method.sync
//
// 付款方式上翻
type AlibabaalihousenewhomepaymentmethodsyncAPIRequest struct {
	model.Params
	// 入参
	_paymentModeReqDto *PaymentModeReqDto
}

// NewAlibabaalihousenewhomepaymentmethodsyncRequest 初始化AlibabaalihousenewhomepaymentmethodsyncAPIRequest对象
func NewAlibabaalihousenewhomepaymentmethodsyncRequest() *AlibabaalihousenewhomepaymentmethodsyncAPIRequest {
	return &AlibabaalihousenewhomepaymentmethodsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomepaymentmethodsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.payment.method.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomepaymentmethodsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomepaymentmethodsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPaymentModeReqDto is PaymentModeReqDto Setter
// 入参
func (r *AlibabaalihousenewhomepaymentmethodsyncAPIRequest) SetPaymentModeReqDto(_paymentModeReqDto *PaymentModeReqDto) error {
	r._paymentModeReqDto = _paymentModeReqDto
	r.Set("payment_mode_req_dto", _paymentModeReqDto)
	return nil
}

// GetPaymentModeReqDto PaymentModeReqDto Getter
func (r AlibabaalihousenewhomepaymentmethodsyncAPIRequest) GetPaymentModeReqDto() *PaymentModeReqDto {
	return r._paymentModeReqDto
}
