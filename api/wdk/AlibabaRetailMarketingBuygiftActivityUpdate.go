package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
更新单品买赠活动 
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新
*/
func AlibabaRetailMarketingBuygiftActivityUpdate(clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivityUpdateRequest, session string) (*wdk.AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse, error) {
    var resp wdk.AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
