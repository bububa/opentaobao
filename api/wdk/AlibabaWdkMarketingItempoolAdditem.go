package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
增加商品池里面的商品 
alibaba.wdk.marketing.itempool.additem

增加商品池里面的商品
*/
func AlibabaWdkMarketingItempoolAdditem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolAdditemAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolAdditemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolAdditemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
