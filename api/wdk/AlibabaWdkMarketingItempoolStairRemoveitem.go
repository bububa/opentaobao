package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
删除换购活动商品 
alibaba.wdk.marketing.itempool.stair.removeitem

删除换购商品
*/
func AlibabaWdkMarketingItempoolStairRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItempoolStairRemoveitemRequest, session string) (*wdk.AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse, error) {
    var resp wdk.AlibabaWdkMarketingItempoolStairRemoveitemAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
