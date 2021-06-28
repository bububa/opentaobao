package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
批量删除买赠商品 
alibaba.wdk.marketing.buygift.item.remove.async

批量删除买赠商品
*/
func AlibabaWdkMarketingBuygiftItemRemoveAsync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingBuygiftItemRemoveAsyncRequest, session string) (*wdk.AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
