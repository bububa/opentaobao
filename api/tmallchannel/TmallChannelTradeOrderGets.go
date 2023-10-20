package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TmallChannelTradeOrderGets 分页查询采购单
// tmall.channel.trade.order.gets
//
// 分页查询采购单
func TmallChannelTradeOrderGets(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeOrderGetsAPIRequest, resp *tmallchannel.TmallChannelTradeOrderGetsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
