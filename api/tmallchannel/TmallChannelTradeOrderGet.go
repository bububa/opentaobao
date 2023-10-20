package tmallchannel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallchannel"
)

// TmallChannelTradeOrderGet 通过主采购单号查询采购单
// tmall.channel.trade.order.get
//
// 通过主采购单号查询采购单
func TmallChannelTradeOrderGet(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeOrderGetAPIRequest, session string) (*tmallchannel.TmallChannelTradeOrderGetAPIResponse, error) {
	var resp tmallchannel.TmallChannelTradeOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
