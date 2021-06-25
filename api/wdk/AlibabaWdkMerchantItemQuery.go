package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
商家商品查询 
alibaba.wdk.merchant.item.query

商家商品查询
*/
func AlibabaWdkMerchantItemQuery(clt *core.SDKClient, req *wdk.AlibabaWdkMerchantItemQueryRequest, session string) (*wdk.AlibabaWdkMerchantItemQueryResponse, error) {
    var resp wdk.AlibabaWdkMerchantItemQueryAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
