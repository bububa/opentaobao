package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressPaymentExchangeGetAPIRequest) Reset() {
	r._checkoutExchangeRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressPaymentExchangeGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.payment.exchange.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressPaymentExchangeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressPaymentExchangeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckoutExchangeRequest is CheckoutExchangeRequest Setter
// 系统自动生成
func (r *AliexpressPaymentExchangeGetAPIRequest) SetCheckoutExchangeRequest(_checkoutExchangeRequest *CheckoutExchangeRequest) error {
	r._checkoutExchangeRequest = _checkoutExchangeRequest
	r.Set("checkout_exchange_request", _checkoutExchangeRequest)
	return nil
}

// GetCheckoutExchangeRequest CheckoutExchangeRequest Getter
func (r AliexpressPaymentExchangeGetAPIRequest) GetCheckoutExchangeRequest() *CheckoutExchangeRequest {
	return r._checkoutExchangeRequest
}

var poolAliexpressPaymentExchangeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressPaymentExchangeGetRequest()
	},
}

// GetAliexpressPaymentExchangeGetRequest 从 sync.Pool 获取 AliexpressPaymentExchangeGetAPIRequest
func GetAliexpressPaymentExchangeGetAPIRequest() *AliexpressPaymentExchangeGetAPIRequest {
	return poolAliexpressPaymentExchangeGetAPIRequest.Get().(*AliexpressPaymentExchangeGetAPIRequest)
}

// ReleaseAliexpressPaymentExchangeGetAPIRequest 将 AliexpressPaymentExchangeGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressPaymentExchangeGetAPIRequest(v *AliexpressPaymentExchangeGetAPIRequest) {
	v.Reset()
	poolAliexpressPaymentExchangeGetAPIRequest.Put(v)
}
