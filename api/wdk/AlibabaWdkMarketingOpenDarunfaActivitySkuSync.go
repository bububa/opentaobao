package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
营销商品数据同步 
alibaba.wdk.marketing.open.darunfa.activity.sku.sync

大润发营销商品数据同步
*/
func AlibabaWdkMarketingOpenDarunfaActivitySkuSync(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenDarunfaActivitySkuSyncRequest, session string) (*wdk.AlibabaWdkMarketingOpenDarunfaActivitySkuSyncResponse, error) {
    var resp wdk.AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
