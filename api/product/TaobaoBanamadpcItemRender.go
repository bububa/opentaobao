package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
新发商品发布页 
taobao.banamadpc.item.render

巴拿马供应商通过此接口新发商品发布页
*/
func TaobaoBanamadpcItemRender(clt *core.SDKClient, req *product.TaobaoBanamadpcItemRenderAPIRequest, session string) (*product.TaobaoBanamadpcItemRenderAPIResponse, error) {
    var resp product.TaobaoBanamadpcItemRenderAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
