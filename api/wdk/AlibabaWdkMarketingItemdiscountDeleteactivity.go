package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
删除商品特价活动 
alibaba.wdk.marketing.itemdiscount.deleteactivity

删除商品特价活动
*/
func AlibabaWdkMarketingItemdiscountDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountDeleteactivityRequest, session string) (*wdk.AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
