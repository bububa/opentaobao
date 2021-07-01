package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
天猫编辑商品规则获取 
tmall.item.update.schema.get

Schema方式编辑天猫商品时，编辑商品规则获取
*/
func TmallItemUpdateSchemaGet(clt *core.SDKClient, req *product.TmallItemUpdateSchemaGetAPIRequest, session string) (*product.TmallItemUpdateSchemaGetAPIResponse, error) {
    var resp product.TmallItemUpdateSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
