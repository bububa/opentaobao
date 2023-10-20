package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresspaymentexchangegetAPIRequest getExchange API请求
// aliexpress.payment.exchange.get
//
// 提供国际汇率服务
type AliexpresspaymentexchangegetAPIRequest struct {
	model.Params
	// 系统自动生成
	_checkoutExchangeRequest *CheckoutExchangeRequest
}

// NewAliexpresspaymentexchangegetRequest 初始化AliexpresspaymentexchangegetAPIRequest对象
func NewAliexpresspaymentexchangegetRequest() *AliexpresspaymentexchangegetAPIRequest {
	return &AliexpresspaymentexchangegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresspaymentexchangegetAPIRequest) GetApiMethodName() string {
	return "aliexpress.payment.exchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresspaymentexchangegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresspaymentexchangegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckoutExchangeRequest is CheckoutExchangeRequest Setter
// 系统自动生成
func (r *AliexpresspaymentexchangegetAPIRequest) SetCheckoutExchangeRequest(_checkoutExchangeRequest *CheckoutExchangeRequest) error {
	r._checkoutExchangeRequest = _checkoutExchangeRequest
	r.Set("checkout_exchange_request", _checkoutExchangeRequest)
	return nil
}

// GetCheckoutExchangeRequest CheckoutExchangeRequest Getter
func (r AliexpresspaymentexchangegetAPIRequest) GetCheckoutExchangeRequest() *CheckoutExchangeRequest {
	return r._checkoutExchangeRequest
}
