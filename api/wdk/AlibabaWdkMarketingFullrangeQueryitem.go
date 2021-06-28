package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
全场活动查询换购品 
alibaba.wdk.marketing.fullrange.queryitem

全场活动查询换购品
*/
func AlibabaWdkMarketingFullrangeQueryitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeQueryitemRequest, session string) (*wdk.AlibabaWdkMarketingFullrangeQueryitemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingFullrangeQueryitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
