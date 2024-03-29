package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) Reset() {
	r._paymentModeReqDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.payment.method.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomePaymentMethodSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomePaymentMethodSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomePaymentMethodSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest
func GetAlibabaAlihouseNewhomePaymentMethodSyncAPIRequest() *AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest {
	return poolAlibabaAlihouseNewhomePaymentMethodSyncAPIRequest.Get().(*AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomePaymentMethodSyncAPIRequest 将 AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomePaymentMethodSyncAPIRequest(v *AlibabaAlihouseNewhomePaymentMethodSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomePaymentMethodSyncAPIRequest.Put(v)
}
