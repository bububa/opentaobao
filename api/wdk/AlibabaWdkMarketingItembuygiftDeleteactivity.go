package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
删除买赠活动 
alibaba.wdk.marketing.itembuygift.deleteactivity

删除买赠活动
*/
func AlibabaWdkMarketingItembuygiftDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
