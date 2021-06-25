package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
全场增加换购品 
alibaba.wdk.marketing.fullrange.addexchangeitem

全场增加换购品
*/
func AlibabaWdkMarketingFullrangeAddexchangeitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeAddexchangeitemRequest, session string) (*wdk.AlibabaWdkMarketingFullrangeAddexchangeitemResponse, error) {
    var resp wdk.AlibabaWdkMarketingFullrangeAddexchangeitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
