package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitpaymentpushAPIRequest 外部推送缴费 API请求
// alibaba.legal.suit.payment.push
//
// 外部推送缴费
type AlibabalegalsuitpaymentpushAPIRequest struct {
	model.Params
	// 推送信息
	_paymentOrderModel *PaymentOrderModel
}

// NewAlibabalegalsuitpaymentpushRequest 初始化AlibabalegalsuitpaymentpushAPIRequest对象
func NewAlibabalegalsuitpaymentpushRequest() *AlibabalegalsuitpaymentpushAPIRequest {
	return &AlibabalegalsuitpaymentpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitpaymentpushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.payment.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitpaymentpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitpaymentpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPaymentOrderModel is PaymentOrderModel Setter
// 推送信息
func (r *AlibabalegalsuitpaymentpushAPIRequest) SetPaymentOrderModel(_paymentOrderModel *PaymentOrderModel) error {
	r._paymentOrderModel = _paymentOrderModel
	r.Set("payment_order_model", _paymentOrderModel)
	return nil
}

// GetPaymentOrderModel PaymentOrderModel Getter
func (r AlibabalegalsuitpaymentpushAPIRequest) GetPaymentOrderModel() *PaymentOrderModel {
	return r._paymentOrderModel
}
