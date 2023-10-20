package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeReturngoodsRefuse 卖家拒绝确认收货
// tmall.exchange.returngoods.refuse
//
// 卖家拒绝买家换货申请
func TmallExchangeReturngoodsRefuse(clt *core.SDKClient, req *exchange.TmallExchangeReturngoodsRefuseAPIRequest, resp *exchange.TmallExchangeReturngoodsRefuseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
