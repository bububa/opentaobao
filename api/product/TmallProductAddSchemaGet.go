package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
产品发布规则获取接口 
tmall.product.add.schema.get

获取用户发布产品的规则
*/
func TmallProductAddSchemaGet(clt *core.SDKClient, req *product.TmallProductAddSchemaGetAPIRequest, session string) (*product.TmallProductAddSchemaGetAPIResponse, error) {
    var resp product.TmallProductAddSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
