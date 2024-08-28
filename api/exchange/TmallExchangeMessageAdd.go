package exchange

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeMessageAdd 卖家创建换货留言
// tmall.exchange.message.add
//
// 卖家创建换货留言
func TmallExchangeMessageAdd(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeMessageAddAPIRequest, resp *exchange.TmallExchangeMessageAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
