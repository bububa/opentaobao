package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeReturngoodsAgree 卖家确认收货
// tmall.exchange.returngoods.agree
//
// 卖家确认收货
func TmallExchangeReturngoodsAgree(clt *core.SDKClient, req *exchange.TmallExchangeReturngoodsAgreeAPIRequest, resp *exchange.TmallExchangeReturngoodsAgreeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
