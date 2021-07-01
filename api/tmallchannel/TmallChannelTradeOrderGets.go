package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

/* TmallChannelTradeOrderGets
分页查询采购单
tmall.channel.trade.order.gets

分页查询采购单 */
func TmallChannelTradeOrderGets(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeOrderGetsAPIRequest, session string) (*tmallchannel.TmallChannelTradeOrderGetsAPIResponse, error) {
	var resp tmallchannel.TmallChannelTradeOrderGetsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
