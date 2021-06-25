package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
更新商品价格 
taobao.item.price.update

更新商品价格
*/
func TaobaoItemPriceUpdate(clt *core.SDKClient, req *product.TaobaoItemPriceUpdateRequest, session string) (*product.TaobaoItemPriceUpdateResponse, error) {
    var resp product.TaobaoItemPriceUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
