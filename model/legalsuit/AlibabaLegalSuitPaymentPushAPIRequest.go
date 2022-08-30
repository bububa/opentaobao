package legalsuit

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitPaymentPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.payment.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitPaymentPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
