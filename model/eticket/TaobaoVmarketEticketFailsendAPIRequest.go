package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoVmarketEticketFailsendAPIRequest
无法发码回调 API请求
taobao.vmarket.eticket.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证 */
type TaobaoVmarketEticketFailsendAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
	// 发码通知时的token
	_token string
	// 错误码
	_errorCode int64
	// 错误信息
	_errorMsg string
}

// New
