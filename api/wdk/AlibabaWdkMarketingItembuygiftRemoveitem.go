package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
移除买赠活动下的商品。【注意，此接口暂不支持并发！】 
alibaba.wdk.marketing.itembuygift.removeitem

移除买赠活动下的商品。【注意，此接口暂不支持并发！】
*/
func AlibabaWdkMarketingItembuygiftRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftRemoveitemAPIRequest, session string) (*wdk.AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItembuygiftRemoveitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
