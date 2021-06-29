package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
供应商停止发货 
tmall.channel.trade.order.stop

供应商停止发货
*/
func TmallChannelTradeOrderStop(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeOrderStopRequest, session string) (*tmallchannel.TmallChannelTradeOrderStopAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeOrderStopAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
