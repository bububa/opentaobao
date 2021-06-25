package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
大商家商品编辑接口 
tmall.item.vip.schema.update

大商家编辑商品
*/
func TmallItemVipSchemaUpdate(clt *core.SDKClient, req *product.TmallItemVipSchemaUpdateRequest, session string) (*product.TmallItemVipSchemaUpdateResponse, error) {
    var resp product.TmallItemVipSchemaUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
