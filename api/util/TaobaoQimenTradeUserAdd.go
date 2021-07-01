package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
添加奇门订单链路用户 
taobao.qimen.trade.user.add

添加奇门订单链路用户
*/
func TaobaoQimenTradeUserAdd(clt *core.SDKClient, req *util.TaobaoQimenTradeUserAddAPIRequest, session string) (*util.TaobaoQimenTradeUserAddAPIResponse, error) {
    var resp util.TaobaoQimenTradeUserAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
