package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-国家标签ID获取 
alibaba.scbp.target.ad.plan.country.id.get

定向推广-国家标签ID获取
*/
func AlibabaScbpTargetAdPlanCountryIdGet(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanCountryIdGetRequest, session string) (*scbp.AlibabaScbpTargetAdPlanCountryIdGetResponse, error) {
    var resp scbp.AlibabaScbpTargetAdPlanCountryIdGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
