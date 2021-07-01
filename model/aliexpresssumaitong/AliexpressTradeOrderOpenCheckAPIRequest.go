package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressTradeOrderOpenCheckAPIRequest
Aliexpress开放平台下单前置检查 API请求
aliexpress.trade.order.open.check

Aliexpress开放平台下单前通过下单入参获取token */
type AliexpressTradeOrderOpenCheckAPIRequest struct {
	model.Params
	// 预下单入参
	_paramPreCreateOrderRequest *PreCreateOrderRequest
	// 客户端信息
	_paramClientInfo *ClientInfo
}

// New
