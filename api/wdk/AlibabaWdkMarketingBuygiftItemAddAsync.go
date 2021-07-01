package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
批量发布买赠商品 
alibaba.wdk.marketing.buygift.item.add.async

批量发布买赠商品
*/
func AlibabaWdkMarketingBuygiftItemAddAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingBuygiftItemAddAsyncAPIRequest, session string) (*wdk.AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingBuygiftItemAddAsyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
