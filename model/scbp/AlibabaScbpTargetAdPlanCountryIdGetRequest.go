package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-国家标签ID获取 API请求
alibaba.scbp.target.ad.plan.country.id.get

定向推广-国家标签ID获取
*/
type AlibabaScbpTargetAdPlanCountryIdGetRequest struct {
    model.Params
}

// 初始化AlibabaScbpTargetAdPlanCountryIdGetRequest对象
func NewAlibabaScbpTargetAdPlanCountryIdGetRequest() *AlibabaScbpTargetAdPlanCountryIdGetRequest{
    return &AlibabaScbpTargetAdPlanCountryIdGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaScbpTargetAdPlanCountryIdGetRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.country.id.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaScbpTargetAdPlanCountryIdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
