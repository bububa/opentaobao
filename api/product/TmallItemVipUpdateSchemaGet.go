package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
vip商家编辑商品的规则获取接口 
tmall.item.vip.update.schema.get

获取vip商家编辑商品的规则
*/
func TmallItemVipUpdateSchemaGet(clt *core.SDKClient, req *product.TmallItemVipUpdateSchemaGetAPIRequest, session string) (*product.TmallItemVipUpdateSchemaGetAPIResponse, error) {
    var resp product.TmallItemVipUpdateSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
