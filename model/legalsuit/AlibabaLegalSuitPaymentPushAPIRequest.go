package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitPaymentPushAPIRequest 外部推送缴费 API请求
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
type AlibabaLegalSuitPaymentPushAPIRequest struct {
	model.Params
	// 推送信息
	_paymentOrderModel *PaymentOrderModel
}

// NewAlibabaLegalSuitPaymentPushRequest 初始化AlibabaLegalSuitPaymentPushAPIRequest对象
func NewAlibabaLegalSuitPaymentPushRequest() *AlibabaLegalSuitPaymentPushAPIRequest {
	return &AlibabaLegalSuitPaymentPushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitPaymentPushAPIRequest) Reset() {
	r._paymentOrderModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitPaymentPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.payment.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitPaymentPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitPaymentPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPaymentOrderModel is PaymentOrderModel Setter
// 推送信息
func (r *AlibabaLegalSuitPaymentPushAPIRequest) SetPaymentOrderModel(_paymentOrderModel *PaymentOrderModel) error {
	r._paymentOrderModel = _paymentOrderModel
	r.Set("payment_order_model", _paymentOrderModel)
	return nil
}

// GetPaymentOrderModel PaymentOrderModel Getter
func (r AlibabaLegalSuitPaymentPushAPIRequest) GetPaymentOrderModel() *PaymentOrderModel {
	return r._paymentOrderModel
}

var poolAlibabaLegalSuitPaymentPushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitPaymentPushRequest()
	},
}

// GetAlibabaLegalSuitPaymentPushRequest 从 sync.Pool 获取 AlibabaLegalSuitPaymentPushAPIRequest
func GetAlibabaLegalSuitPaymentPushAPIRequest() *AlibabaLegalSuitPaymentPushAPIRequest {
	return poolAlibabaLegalSuitPaymentPushAPIRequest.Get().(*AlibabaLegalSuitPaymentPushAPIRequest)
}

// ReleaseAlibabaLegalSuitPaymentPushAPIRequest 将 AlibabaLegalSuitPaymentPushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitPaymentPushAPIRequest(v *AlibabaLegalSuitPaymentPushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitPaymentPushAPIRequest.Put(v)
}
