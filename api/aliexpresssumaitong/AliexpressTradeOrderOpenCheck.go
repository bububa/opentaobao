package aliexpresssumaitong

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// AliexpressTradeOrderOpenCheck Aliexpress开放平台下单前置检查
// aliexpress.trade.order.open.check
//
// Aliexpress开放平台下单前通过下单入参获取token
func AliexpressTradeOrderOpenCheck(clt *core.SDKClient, req *aliexpresssumaitong.AliexpressTradeOrderOpenCheckAPIRequest, resp *aliexpresssumaitong.AliexpressTradeOrderOpenCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
