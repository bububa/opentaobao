package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
全场活动删除活动接口 
alibaba.wdk.marketing.fullrange.deleteactivity

全场活动删除活动
*/
func AlibabaWdkMarketingFullrangeDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
