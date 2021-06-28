package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
特价批量移除商品 
alibaba.wdk.marketing.discount.item.remove.async

特价批量移除商品
*/
func AlibabaWdkMarketingDiscountItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingDiscountItemRemoveAsyncRequest, session string) (*wdk.AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
