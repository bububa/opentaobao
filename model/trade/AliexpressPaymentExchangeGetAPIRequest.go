package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressPaymentExchangeGetAPIRequest
getExchange API请求
aliexpress.payment.exchange.get

提供国际汇率服务 */
type AliexpressPaymentExchangeGetAPIRequest struct {
	model.Params
	// 系统自动生成
	_checkoutExchangeRequest *CheckoutExchangeRequest
}

// New
