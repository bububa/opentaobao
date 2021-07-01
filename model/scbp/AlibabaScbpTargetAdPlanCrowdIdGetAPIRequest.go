package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-人群标签ID获取(店铺老客、优选人群) API请求
alibaba.scbp.target.ad.plan.crowd.id.get

定向推广-人群标签ID获取(店铺老客、优选人群)
*/
type AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest struct {
    model.Params
}

// 初始化AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest对象
func NewAlibabaScbpTargetAdPlanCrowdIdGetRequest() *AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest{
    return &AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.crowd.id.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
