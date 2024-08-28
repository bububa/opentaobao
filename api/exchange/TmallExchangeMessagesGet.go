package exchange

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeMessagesGet 查询换货订单留言列表
// tmall.exchange.messages.get
//
// 查询换货订单留言列表
func TmallExchangeMessagesGet(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeMessagesGetAPIRequest, resp *exchange.TmallExchangeMessagesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
