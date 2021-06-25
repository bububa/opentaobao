package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取子属性 
taobao.banamadpc.item.select.prop

巴拿马供应商通过此接口获取子属性
*/
func TaobaoBanamadpcItemSelectProp(clt *core.SDKClient, req *product.TaobaoBanamadpcItemSelectPropRequest, session string) (*product.TaobaoBanamadpcItemSelectPropResponse, error) {
    var resp product.TaobaoBanamadpcItemSelectPropAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
