package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查找特价活动 
alibaba.wdk.marketing.itemdiscount.queryactivity

查找特价活动
*/
func AlibabaWdkMarketingItemdiscountQueryactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountQueryactivityRequest, session string) (*wdk.AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
