package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
版本数量查询 
alibaba.wdk.marketing.open.version.count

版本数量查询
*/
func AlibabaWdkMarketingOpenVersionCount(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingOpenVersionCountAPIRequest, session string) (*wdk.AlibabaWdkMarketingOpenVersionCountAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingOpenVersionCountAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
