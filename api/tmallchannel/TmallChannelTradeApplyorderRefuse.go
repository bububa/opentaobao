package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
供应商审核拒绝采购申请单 
tmall.channel.trade.applyorder.refuse

供应商审核拒绝采购申请单
*/
func TmallChannelTradeApplyorderRefuse(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeApplyorderRefuseRequest, session string) (*tmallchannel.TmallChannelTradeApplyorderRefuseAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeApplyorderRefuseAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
