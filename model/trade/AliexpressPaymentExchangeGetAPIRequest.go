package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressPaymentExchangeGetAPIRequest getExchange API请求
// aliexpress.payment.exchange.get
//
// 提供国际汇率服务
type AliexpressPaymentExchangeGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_checkoutExchangeRequest *CheckoutExchangeRequest
}

// NewAliexpressPaymentExchangeGetRequest 初始化AliexpressPaymentExchangeGetAPIRequest对象
func NewAliexpressPaymentExchangeGetRequest() *AliexpressPaymentExchangeGetAPIRequest {
	return &AliexpressPaymentExchangeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressPaymentExchangeGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.payment.exchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressPaymentExchangeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CheckoutExchangeRequest Setter
// 系统自动生成
func (r *AliexpressPaymentExchangeGetAPIRequest) SetCheckoutExchangeRequest(_checkoutExchangeRequest *CheckoutExchangeRequest) error {
	r._checkoutExchangeRequest = _checkoutExchangeRequest
	r.Set("checkout_exchange_request", _checkoutExchangeRequest)
	return nil
}

// Get CheckoutExchangeRequest Getter
func (r AliexpressPaymentExchangeGetAPIRequest) GetCheckoutExchangeRequest() *CheckoutExchangeRequest {
	return r._checkoutExchangeRequest
}
