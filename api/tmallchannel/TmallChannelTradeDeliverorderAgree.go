package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
供应商审核通过发货确认 
tmall.channel.trade.deliverorder.agree

供应商通过收货确认单
*/
func TmallChannelTradeDeliverorderAgree(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeDeliverorderAgreeAPIRequest, session string) (*tmallchannel.TmallChannelTradeDeliverorderAgreeAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeDeliverorderAgreeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
