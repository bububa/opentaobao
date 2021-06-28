package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取匹配产品规则 
tmall.product.match.schema.get

ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
*/
func TmallProductMatchSchemaGet(clt *core.SDKClient, req *product.TmallProductMatchSchemaGetRequest, session string) (*product.TmallProductMatchSchemaGetAPIResponse, error) {
    var resp product.TmallProductMatchSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
