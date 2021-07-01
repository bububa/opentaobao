package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
全场活动删除购品 
alibaba.wdk.marketing.fullrange.removeitem

删除换购商品
*/
func AlibabaWdkMarketingFullrangeRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeRemoveitemAPIRequest, session string) (*wdk.AlibabaWdkMarketingFullrangeRemoveitemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingFullrangeRemoveitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
