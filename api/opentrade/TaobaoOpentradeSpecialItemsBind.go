package opentrade

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/opentrade"
)

/* 
专属下单场景商品绑定 
taobao.opentrade.special.items.bind

专属下单场景商品绑定
*/
func TaobaoOpentradeSpecialItemsBind(clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialItemsBindAPIRequest, session string) (*opentrade.TaobaoOpentradeSpecialItemsBindAPIResponse, error) {
    var resp opentrade.TaobaoOpentradeSpecialItemsBindAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
