package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-新建计划 APIRequest
alibaba.scbp.target.ad.plan.add

定向推广-新建单条计划
*/
type AlibabaScbpTargetAdPlanAddRequest struct {
    model.Params

    // 定向推广基础信息
    topP4pBasicQuickCampaign   *BasicQuickCampaign 

}

func NewAlibabaScbpTargetAdPlanAddRequest() *AlibabaScbpTargetAdPlanAddRequest{
    return &AlibabaScbpTargetAdPlanAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanAddRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.add"
}

func (r AlibabaScbpTargetAdPlanAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanAddRequest) SetTopP4pBasicQuickCampaign(topP4pBasicQuickCampaign *BasicQuickCampaign) error {
    r.topP4pBasicQuickCampaign = topP4pBasicQuickCampaign
    r.Set("top_p4p_basic_quick_campaign", topP4pBasicQuickCampaign)
    return nil
}

func (r AlibabaScbpTargetAdPlanAddRequest) GetTopP4pBasicQuickCampaign() *BasicQuickCampaign {
    return r.topP4pBasicQuickCampaign
}

