package scbp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-更新推广计划的基础信息 APIRequest
alibaba.scbp.target.ad.plan.update

定向推广-更新推广计划的基础信息
*/
type AlibabaScbpTargetAdPlanUpdateRequest struct {
    model.Params

    // TopP4pBasicQuickCampaign
    topP4pBasicQuickCampaign   *TopP4pBasicQuickCampaign 

}

func NewAlibabaScbpTargetAdPlanUpdateRequest() *AlibabaScbpTargetAdPlanUpdateRequest{
    return &AlibabaScbpTargetAdPlanUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaScbpTargetAdPlanUpdateRequest) GetApiMethodName() string {
    return "alibaba.scbp.target.ad.plan.update"
}

func (r AlibabaScbpTargetAdPlanUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaScbpTargetAdPlanUpdateRequest) SetTopP4pBasicQuickCampaign(topP4pBasicQuickCampaign *TopP4pBasicQuickCampaign) error {
    r.topP4pBasicQuickCampaign = topP4pBasicQuickCampaign
    r.Set("top_p4p_basic_quick_campaign", topP4pBasicQuickCampaign)
    return nil
}

func (r AlibabaScbpTargetAdPlanUpdateRequest) GetTopP4pBasicQuickCampaign() *TopP4pBasicQuickCampaign {
    return r.topP4pBasicQuickCampaign
}

