package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
增加买赠活动商品。【注意，此接口暂不支持并发！】 
alibaba.wdk.marketing.itembuygift.additem

增加买赠活动商品。【注意，此接口暂不支持并发！】
*/
func AlibabaWdkMarketingItembuygiftAdditem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftAdditemAPIRequest, session string) (*wdk.AlibabaWdkMarketingItembuygiftAdditemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItembuygiftAdditemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
