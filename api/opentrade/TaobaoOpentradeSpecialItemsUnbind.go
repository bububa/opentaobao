package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
专属下单场景商品解绑 
taobao.opentrade.special.items.unbind

专属下单场景商品解绑
*/
func TaobaoOpentradeSpecialItemsUnbind(clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialItemsUnbindRequest, session string) (*opentrade.TaobaoOpentradeSpecialItemsUnbindAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeSpecialItemsUnbindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
