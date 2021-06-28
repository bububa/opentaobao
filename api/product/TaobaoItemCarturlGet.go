package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
加购URL获取 
taobao.item.carturl.get

获取加购URL，支持添加商品到购物车
*/
func TaobaoItemCarturlGet(clt *core.SDKClient, req *product.TaobaoItemCarturlGetRequest, session string) (*product.TaobaoItemCarturlGetAPIResponse, error) {
    var resp product.TaobaoItemCarturlGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
