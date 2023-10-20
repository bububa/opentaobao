package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeMessageAdd 卖家创建换货留言
// tmall.exchange.message.add
//
// 卖家创建换货留言
func TmallExchangeMessageAdd(clt *core.SDKClient, req *exchange.TmallExchangeMessageAddAPIRequest, resp *exchange.TmallExchangeMessageAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
