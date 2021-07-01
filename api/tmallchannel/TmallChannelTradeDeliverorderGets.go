package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

/* TmallChannelTradeDeliverorderGets
查询发货单列表
tmall.channel.trade.deliverorder.gets

查询发货单列表 */
func TmallChannelTradeDeliverorderGets(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeDeliverorderGetsAPIRequest, session string) (*tmallchannel.TmallChannelTradeDeliverorderGetsAPIResponse, error) {
	var resp tmallchannel.TmallChannelTradeDeliverorderGetsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
