package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
创建商品池活动【同城零售】 
alibaba.retail.marketing.itempool.activity.create

同城零售商品池活动创建
*/
func AlibabaRetailMarketingItempoolActivityCreate(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolActivityCreateRequest, session string) (*wdk.AlibabaRetailMarketingItempoolActivityCreateResponse, error) {
    var resp wdk.AlibabaRetailMarketingItempoolActivityCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
