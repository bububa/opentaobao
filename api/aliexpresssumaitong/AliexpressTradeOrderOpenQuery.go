package aliexpresssumaitong

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// AliexpressTradeOrderOpenQuery Aliexpress开放平台订单查询
// aliexpress.trade.order.open.query
//
// Aliexpress开放平台订单信息查询
func AliexpressTradeOrderOpenQuery(clt *core.SDKClient, req *aliexpresssumaitong.AliexpressTradeOrderOpenQueryAPIRequest, resp *aliexpresssumaitong.AliexpressTradeOrderOpenQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
