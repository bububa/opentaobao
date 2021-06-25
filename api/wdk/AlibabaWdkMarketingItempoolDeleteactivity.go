package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
删除商品池活动 
alibaba.wdk.marketing.itempool.deleteactivity

删除商品池活动
*/
func AlibabaWdkMarketingItempoolDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolDeleteactivityRequest, session string) (*wdk.AlibabaWdkMarketingItempoolDeleteactivityResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolDeleteactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
