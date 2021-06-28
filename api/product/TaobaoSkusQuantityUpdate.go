package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
SKU库存修改 
taobao.skus.quantity.update

提供按照全量/增量的方式批量修改SKU库存的功能
*/
func TaobaoSkusQuantityUpdate(clt *core.SDKClient, req *product.TaobaoSkusQuantityUpdateRequest, session string) (*product.TaobaoSkusQuantityUpdateAPIResponse, error) {
    var resp product.TaobaoSkusQuantityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
