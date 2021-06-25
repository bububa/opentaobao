package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
更新单品特价活动【同城零售】 
alibaba.retail.marketing.itemdiscount.activity.update

同城零售单品特价活动更新
*/
func AlibabaRetailMarketingItemdiscountActivityUpdate(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivityUpdateRequest, session string) (*wdk.AlibabaRetailMarketingItemdiscountActivityUpdateResponse, error) {
    var resp wdk.AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
