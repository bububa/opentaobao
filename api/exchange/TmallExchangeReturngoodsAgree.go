package exchange

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeReturngoodsAgree 卖家确认收货
// tmall.exchange.returngoods.agree
//
// 卖家确认收货
func TmallExchangeReturngoodsAgree(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeReturngoodsAgreeAPIRequest, resp *exchange.TmallExchangeReturngoodsAgreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
