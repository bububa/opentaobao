package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
查询买赠活动下的商品 
alibaba.wdk.marketing.itembuygift.queryitems

查询买赠活动下的商品
*/
func AlibabaWdkMarketingItembuygiftQueryitems(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftQueryitemsAPIRequest, session string) (*wdk.AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItembuygiftQueryitemsAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
