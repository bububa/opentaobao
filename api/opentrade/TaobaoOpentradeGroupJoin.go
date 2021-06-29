package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
组团购场景参团 
taobao.opentrade.group.join

组团购场景下，用户参团
*/
func TaobaoOpentradeGroupJoin(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupJoinRequest, session string) (*opentrade.TaobaoOpentradeGroupJoinAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeGroupJoinAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
