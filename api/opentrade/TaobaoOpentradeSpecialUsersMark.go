package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
专属下单可购买用户标记 
taobao.opentrade.special.users.mark

对于专属下单的交易场景，根据openid标记用户可购买商品
*/
func TaobaoOpentradeSpecialUsersMark(clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialUsersMarkRequest, session string) (*opentrade.TaobaoOpentradeSpecialUsersMarkAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeSpecialUsersMarkAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
