package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
供应商审核同意采购申请单 
tmall.channel.trade.applyorder.agree

供应商审核同意采购申请单
*/
func TmallChannelTradeApplyorderAgree(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeApplyorderAgreeRequest, session string) (*tmallchannel.TmallChannelTradeApplyorderAgreeAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeApplyorderAgreeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
