package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
定向推广-人群标签ID获取(店铺老客、优选人群) 
alibaba.scbp.target.ad.plan.crowd.id.get

定向推广-人群标签ID获取(店铺老客、优选人群)
*/
func AlibabaScbpTargetAdPlanCrowdIdGet(clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest, session string) (*scbp.AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse, error) {
    var resp scbp.AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
