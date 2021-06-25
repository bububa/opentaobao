package product

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/product"
)

/* 
商品发布 
alibaba.item.publish.submit

新商品发布，提交商品发布信息
*/
func AlibabaItemPublishSubmit(clt *core.SDKClient, req *product.AlibabaItemPublishSubmitRequest, session string) (*product.AlibabaItemPublishSubmitResponse, error) {
    var resp product.AlibabaItemPublishSubmitAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
