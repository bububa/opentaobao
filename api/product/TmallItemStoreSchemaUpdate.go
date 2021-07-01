package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
天猫门店商品编辑 
tmall.item.store.schema.update

天猫门店商品编辑
*/
func TmallItemStoreSchemaUpdate(clt *core.SDKClient, req *product.TmallItemStoreSchemaUpdateAPIRequest, session string) (*product.TmallItemStoreSchemaUpdateAPIResponse, error) {
    var resp product.TmallItemStoreSchemaUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
