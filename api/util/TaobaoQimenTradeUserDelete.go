package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
删除奇门订单链路用户 
taobao.qimen.trade.user.delete

删除奇门订单链路用户
*/
func TaobaoQimenTradeUserDelete(clt *core.SDKClient, req *util.TaobaoQimenTradeUserDeleteAPIRequest, session string) (*util.TaobaoQimenTradeUserDeleteAPIResponse, error) {
    var resp util.TaobaoQimenTradeUserDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
