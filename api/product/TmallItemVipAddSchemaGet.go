package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
vip商家发布商品的获取规则接口 
tmall.item.vip.add.schema.get

获取vip商家发布商品的规则
*/
func TmallItemVipAddSchemaGet(clt *core.SDKClient, req *product.TmallItemVipAddSchemaGetAPIRequest, session string) (*product.TmallItemVipAddSchemaGetAPIResponse, error) {
    var resp product.TmallItemVipAddSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
