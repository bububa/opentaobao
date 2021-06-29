package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
尖货交易可购买用户标记 
taobao.opentrade.queue.users.mark

尖货交易用户标记信息回传，根据openId标记用户可购买商品
*/
func TaobaoOpentradeQueueUsersMark(clt *core.SDKClient, req *opentrade.TaobaoOpentradeQueueUsersMarkRequest, session string) (*opentrade.TaobaoOpentradeQueueUsersMarkAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeQueueUsersMarkAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
