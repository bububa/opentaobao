package tmallchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallchannel"
)

/* 
获取采购申请单列表 
tmall.channel.trade.applyorder.gets

分页查询采购申请单列表
*/
func TmallChannelTradeApplyorderGets(clt *core.SDKClient, req *tmallchannel.TmallChannelTradeApplyorderGetsRequest, session string) (*tmallchannel.TmallChannelTradeApplyorderGetsAPIResponse, error) {
    var resp tmallchannel.TmallChannelTradeApplyorderGetsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
