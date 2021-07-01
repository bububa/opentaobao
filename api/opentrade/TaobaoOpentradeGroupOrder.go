package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
组团购获取订单列表 
taobao.opentrade.group.order

组团购场景下，获取开团的订单列表
*/
func TaobaoOpentradeGroupOrder(clt *core.SDKClient, req *opentrade.TaobaoOpentradeGroupOrderAPIRequest, session string) (*opentrade.TaobaoOpentradeGroupOrderAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeGroupOrderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
