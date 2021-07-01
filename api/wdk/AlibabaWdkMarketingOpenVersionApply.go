package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
数据同步版本号申请 
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请
*/
func AlibabaWdkMarketingOpenVersionApply(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenVersionApplyAPIRequest, session string) (*wdk.AlibabaWdkMarketingOpenVersionApplyAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingOpenVersionApplyAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
