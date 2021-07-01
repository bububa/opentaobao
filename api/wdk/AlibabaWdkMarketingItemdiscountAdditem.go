package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
报名特价商品 
alibaba.wdk.marketing.itemdiscount.additem

在商品特价活动中报名特价商品
*/
func AlibabaWdkMarketingItemdiscountAdditem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountAdditemAPIRequest, session string) (*wdk.AlibabaWdkMarketingItemdiscountAdditemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItemdiscountAdditemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
