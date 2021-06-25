package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
天猫门店商品修改规则获取 
tmall.item.store.update.schema.get

天猫门店商品修改规则获取
*/
func TmallItemStoreUpdateSchemaGet(clt *core.SDKClient, req *product.TmallItemStoreUpdateSchemaGetRequest, session string) (*product.TmallItemStoreUpdateSchemaGetResponse, error) {
    var resp product.TmallItemStoreUpdateSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
