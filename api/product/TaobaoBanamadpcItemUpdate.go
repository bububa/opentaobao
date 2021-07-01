package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
编辑商品 
taobao.banamadpc.item.update

巴拿马供应商通过此接口编辑商品
*/
func TaobaoBanamadpcItemUpdate(clt *core.SDKClient, req *product.TaobaoBanamadpcItemUpdateAPIRequest, session string) (*product.TaobaoBanamadpcItemUpdateAPIResponse, error) {
    var resp product.TaobaoBanamadpcItemUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
