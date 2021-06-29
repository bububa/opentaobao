package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
创建渠道分销单 
tmall.channel.trade.order.create

创建渠道分销单
*/
func TmallChannelTradeOrderCreate(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeOrderCreateRequest, session string) (*tmallchannel.TmallChannelTradeOrderCreateAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeOrderCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
