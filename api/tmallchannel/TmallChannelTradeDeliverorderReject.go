package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

/* TmallChannelTradeDeliverorderReject
供应商拒绝收货确认单
tmall.channel.trade.deliverorder.reject

供应商拒绝收货确认单 */
func TmallChannelTradeDeliverorderReject(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeDeliverorderRejectAPIRequest, session string) (*tmallchannel.TmallChannelTradeDeliverorderRejectAPIResponse, error) {
	var resp tmallchannel.TmallChannelTradeDeliverorderRejectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
