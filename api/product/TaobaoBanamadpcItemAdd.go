package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
新发商品 
taobao.banamadpc.item.add

巴拿马供应商通过此接口新发商品
*/
func TaobaoBanamadpcItemAdd(clt *core.SDKClient, req *product.TaobaoBanamadpcItemAddRequest, session string) (*product.TaobaoBanamadpcItemAddAPIResponse, error) {
    var resp product.TaobaoBanamadpcItemAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
