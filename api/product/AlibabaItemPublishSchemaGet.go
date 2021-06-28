package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
获取商品发布规则信息 
alibaba.item.publish.schema.get

新商品发布，获取商品发布规则信息
*/
func AlibabaItemPublishSchemaGet(clt *core.SDKClient, req *product.AlibabaItemPublishSchemaGetRequest, session string) (*product.AlibabaItemPublishSchemaGetAPIResponse, error) {
    var resp product.AlibabaItemPublishSchemaGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
