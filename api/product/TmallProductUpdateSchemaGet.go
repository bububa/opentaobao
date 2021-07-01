package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
产品更新规则获取接口 
tmall.product.update.schema.get

获取用户更新产品的规则
*/
func TmallProductUpdateSchemaGet(clt *core.SDKClient, req *product.TmallProductUpdateSchemaGetAPIRequest, session string) (*product.TmallProductUpdateSchemaGetAPIResponse, error) {
    var resp product.TmallProductUpdateSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
