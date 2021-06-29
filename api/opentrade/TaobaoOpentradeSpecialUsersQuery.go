package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
专属下单标记信息查询 
taobao.opentrade.special.users.query

专属下单标记信息查询
*/
func TaobaoOpentradeSpecialUsersQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialUsersQueryRequest, session string) (*opentrade.TaobaoOpentradeSpecialUsersQueryAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeSpecialUsersQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
