package util

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/util"
)

/* 
获取奇门用户列表 
taobao.qimen.trade.users.get

获取已开通奇门订单服务的用户列表
*/
func TaobaoQimenTradeUsersGet(clt *core.SDKClient, req *util.TaobaoQimenTradeUsersGetAPIRequest, session string) (*util.TaobaoQimenTradeUsersGetAPIResponse, error) {
    var resp util.TaobaoQimenTradeUsersGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
