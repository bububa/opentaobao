package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
移除报名的商品 
alibaba.wdk.marketing.itemdiscount.removeitem

将报名特价活动的商品从特价活动中移除
*/
func AlibabaWdkMarketingItemdiscountRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountRemoveitemRequest, session string) (*wdk.AlibabaWdkMarketingItemdiscountRemoveitemResponse, error) {
    var resp wdk.AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
