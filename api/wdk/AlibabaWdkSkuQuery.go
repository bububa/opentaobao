package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询商品 
alibaba.wdk.sku.query

查询商品
*/
func AlibabaWdkSkuQuery(clt *core.SDKClient, req *wdk.AlibabaWdkSkuQueryRequest, session string) (*wdk.AlibabaWdkSkuQueryResponse, error) {
    var resp wdk.AlibabaWdkSkuQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
