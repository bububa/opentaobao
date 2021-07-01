package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
换购商品查询 
alibaba.wdk.marketing.itempool.stair.queryitem

换购商品查询
*/
func AlibabaWdkMarketingItempoolStairQueryitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolStairQueryitemAPIRequest, session string) (*wdk.AlibabaWdkMarketingItempoolStairQueryitemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolStairQueryitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
