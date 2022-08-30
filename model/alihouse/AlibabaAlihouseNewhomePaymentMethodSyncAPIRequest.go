package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest 付款方式上翻 API请求
// alibaba.alihouse.newhome.payment.method.sync
//
// 付款方式上翻
type AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest struct {
	model.Params
	// 入参
	_paymentModeReqDto *PaymentModeReqDto
}

// NewAlibabaAlihouseNewhomePaymentMethodSyncRequest 初始化AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest对象
func NewAlibabaAlihouseNewhomePaymentMethodSyncRequest() *AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest {
	return &AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.payment.method.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPaymentModeReqDto is PaymentModeReqDto Setter
// 入参
func (r *AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) SetPaymentModeReqDto(_paymentModeReqDto *PaymentModeReqDto) error {
	r._paymentModeReqDto = _paymentModeReqDto
	r.Set("payment_mode_req_dto", _paymentModeReqDto)
	return nil
}

// GetPaymentModeReqDto PaymentModeReqDto Getter
func (r AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) GetPaymentModeReqDto() *PaymentModeReqDto {
	return r._paymentModeReqDto
}
