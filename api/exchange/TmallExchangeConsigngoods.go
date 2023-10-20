package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeConsigngoods 卖家发货
// tmall.exchange.consigngoods
//
// 卖家发货
func TmallExchangeConsigngoods(clt *core.SDKClient, req *exchange.TmallExchangeConsigngoodsAPIRequest, resp *exchange.TmallExchangeConsigngoodsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
