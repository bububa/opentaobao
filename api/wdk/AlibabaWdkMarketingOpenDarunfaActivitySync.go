package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
活动数据同步 
alibaba.wdk.marketing.open.darunfa.activity.sync

大润发活动数据同步
*/
func AlibabaWdkMarketingOpenDarunfaActivitySync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenDarunfaActivitySyncRequest, session string) (*wdk.AlibabaWdkMarketingOpenDarunfaActivitySyncResponse, error) {
    var resp wdk.AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
