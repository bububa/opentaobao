package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
特价批量新增商品 
alibaba.wdk.marketing.discount.item.add.async

新分组模型下新增商品
*/
func AlibabaWdkMarketingDiscountItemAddAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingDiscountItemAddAsyncRequest, session string) (*wdk.AlibabaWdkMarketingDiscountItemAddAsyncResponse, error) {
    var resp wdk.AlibabaWdkMarketingDiscountItemAddAsyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
