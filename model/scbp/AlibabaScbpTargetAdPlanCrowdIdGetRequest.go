package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-人群标签ID获取(店铺老客、优选人群) APIRequest
alibaba.scbp.target.ad.plan.crowd.id.get

定向推广-人群标签ID获取(店铺老客、优选人群)
*/
type AlibabaScbpTargetAdPlanCrowdIdGetRequest struct {
    model.Params

}

func NewAlibabaScbpTargetAdPlanCrowdIdGetRequest() *AlibabaScbpTargetAdPlanCrowdIdGetRequest{
    return &AlibabaScbpTargetAdPlanCrowdIdGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanCrowdIdGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.crowd.id.get"
}

func (r AlibabaScbpTargetAdPlanCrowdIdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


