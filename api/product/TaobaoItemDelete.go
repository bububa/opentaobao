package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
删除单条商品 
taobao.item.delete

删除单条商品
*/
func TaobaoItemDelete(clt *core.SDKClient, req *product.TaobaoItemDeleteRequest, session string) (*product.TaobaoItemDeleteAPIResponse, error) {
    var resp product.TaobaoItemDeleteAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
