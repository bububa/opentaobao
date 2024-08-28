package tmallchannel

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TmallChannelTradeOrderGet 通过主采购单号查询采购单
// tmall.channel.trade.order.get
//
// 通过主采购单号查询采购单
func TmallChannelTradeOrderGet(ctx context.Context, clt *core.SDKClient, req *tmallchannel.TmallChannelTradeOrderGetAPIRequest, resp *tmallchannel.TmallChannelTradeOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
