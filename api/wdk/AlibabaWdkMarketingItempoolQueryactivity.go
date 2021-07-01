package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查找商品池活动 
alibaba.wdk.marketing.itempool.queryactivity

查找商品池活动
*/
func AlibabaWdkMarketingItempoolQueryactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolQueryactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolQueryactivityAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolQueryactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
