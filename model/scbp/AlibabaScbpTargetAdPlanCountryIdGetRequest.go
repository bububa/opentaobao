package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-国家标签ID获取 APIRequest
alibaba.scbp.target.ad.plan.country.id.get

定向推广-国家标签ID获取
*/
type AlibabaScbpTargetAdPlanCountryIdGetRequest struct {
    model.Params

}

func NewAlibabaScbpTargetAdPlanCountryIdGetRequest() *AlibabaScbpTargetAdPlanCountryIdGetRequest{
    return &AlibabaScbpTargetAdPlanCountryIdGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanCountryIdGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.country.id.get"
}

func (r AlibabaScbpTargetAdPlanCountryIdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


