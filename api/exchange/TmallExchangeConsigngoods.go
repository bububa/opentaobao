package exchange

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeConsigngoods 卖家发货
// tmall.exchange.consigngoods
//
// 卖家发货
func TmallExchangeConsigngoods(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeConsigngoodsAPIRequest, resp *exchange.TmallExchangeConsigngoodsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
