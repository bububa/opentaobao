package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
通过发货单单号获取发货单的详情 
tmall.channel.trade.deliverorder.get

通过发货单单号获取发货单的详情
*/
func TmallChannelTradeDeliverorderGet(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeDeliverorderGetAPIRequest, session string) (*tmallchannel.TmallChannelTradeDeliverorderGetAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeDeliverorderGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
