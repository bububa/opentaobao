package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
大商家商品发布接口 
tmall.item.vip.schema.add

大商家商品发布接口
*/
func TmallItemVipSchemaAdd(clt *core.SDKClient, req *product.TmallItemVipSchemaAddAPIRequest, session string) (*product.TmallItemVipSchemaAddAPIResponse, error) {
    var resp product.TmallItemVipSchemaAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
