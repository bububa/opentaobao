package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
供应商查询退款单 
tmall.channel.trade.refundorder.gets

供应商分页查询退款单
*/
func TmallChannelTradeRefundorderGets(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeRefundorderGetsAPIRequest, session string) (*tmallchannel.TmallChannelTradeRefundorderGetsAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeRefundorderGetsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
